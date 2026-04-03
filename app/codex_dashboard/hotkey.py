from __future__ import annotations

import ctypes
import queue
import threading
from ctypes import wintypes

user32 = ctypes.WinDLL("user32", use_last_error=True)
kernel32 = ctypes.WinDLL("kernel32", use_last_error=True)

WM_HOTKEY = 0x0312
WM_QUIT = 0x0012

MOD_ALT = 0x0001
MOD_CONTROL = 0x0002
MOD_SHIFT = 0x0004
MOD_WIN = 0x0008


class POINT(ctypes.Structure):
    _fields_ = [
        ("x", wintypes.LONG),
        ("y", wintypes.LONG),
    ]


class MSG(ctypes.Structure):
    _fields_ = [
        ("hwnd", wintypes.HWND),
        ("message", wintypes.UINT),
        ("wParam", wintypes.WPARAM),
        ("lParam", wintypes.LPARAM),
        ("time", wintypes.DWORD),
        ("pt", POINT),
    ]


def parse_hotkey(hotkey: str) -> tuple[int, int]:
    modifiers = 0
    virtual_key = None
    for part in [segment.strip() for segment in hotkey.split("+") if segment.strip()]:
        lower = part.lower()
        if lower == "ctrl":
            modifiers |= MOD_CONTROL
            continue
        if lower == "alt":
            modifiers |= MOD_ALT
            continue
        if lower == "shift":
            modifiers |= MOD_SHIFT
            continue
        if lower in {"win", "meta"}:
            modifiers |= MOD_WIN
            continue
        if lower == "space":
            virtual_key = 0x20
            continue
        if lower.startswith("f") and lower[1:].isdigit():
            number = int(lower[1:])
            if 1 <= number <= 24:
                virtual_key = 0x6F + number
                continue
        if len(part) == 1:
            virtual_key = ord(part.upper())
            continue
        raise ValueError(f"Unsupported hotkey segment: {part}")
    if modifiers == 0 or virtual_key is None:
        raise ValueError(f"Hotkey must include modifiers and a key: {hotkey}")
    return modifiers, virtual_key


class GlobalHotkey:
    def __init__(self, hotkey: str, callback, hotkey_id: int = 1) -> None:
        self.hotkey = hotkey
        self.callback = callback
        self.hotkey_id = hotkey_id
        self.modifiers, self.virtual_key = parse_hotkey(hotkey)
        self._registered = False
        self._thread: threading.Thread | None = None
        self._thread_id: int | None = None
        self._register_event = threading.Event()
        self._register_error: Exception | None = None
        self._pending_callbacks: queue.SimpleQueue[None] = queue.SimpleQueue()

    def register(self) -> None:
        if self._registered:
            return
        self._register_error = None
        self._register_event.clear()
        self._thread = threading.Thread(
            target=self._message_loop,
            name=f"CodexDashboardHotkey-{self.hotkey_id}",
            daemon=True,
        )
        self._thread.start()
        if not self._register_event.wait(timeout=2.0):
            raise TimeoutError(f"Timed out while registering hotkey {self.hotkey}")
        if self._register_error is not None:
            error = self._register_error
            self._register_error = None
            self._thread = None
            raise error

    def _message_loop(self) -> None:
        try:
            self._thread_id = int(kernel32.GetCurrentThreadId())
            if not user32.RegisterHotKey(
                None,
                self.hotkey_id,
                self.modifiers,
                self.virtual_key,
            ):
                error = ctypes.get_last_error()
                self._register_error = OSError(error, f"RegisterHotKey failed for {self.hotkey}")
                return
            self._registered = True
            self._register_event.set()

            message = MSG()
            while True:
                result = user32.GetMessageW(ctypes.byref(message), None, 0, 0)
                if result == -1:
                    error = ctypes.get_last_error()
                    self._register_error = OSError(error, f"GetMessageW failed for {self.hotkey}")
                    break
                if result == 0:
                    break
                if message.message == WM_HOTKEY and int(message.wParam) == self.hotkey_id:
                    self._pending_callbacks.put(None)
        except Exception as exc:  # pragma: no cover - defensive threading fallback
            self._register_error = exc
        finally:
            if not self._register_event.is_set():
                self._register_event.set()
            if self._registered:
                user32.UnregisterHotKey(None, self.hotkey_id)
                self._registered = False
            self._thread_id = None

    def poll(self) -> None:
        while True:
            try:
                self._pending_callbacks.get_nowait()
            except queue.Empty:
                break
            self.callback()

    def unregister(self) -> None:
        thread_id = self._thread_id
        thread = self._thread
        if thread_id is not None:
            user32.PostThreadMessageW(thread_id, WM_QUIT, 0, 0)
        if thread is not None:
            thread.join(timeout=2.0)
        self._thread = None
        self._thread_id = None
        self._registered = False
