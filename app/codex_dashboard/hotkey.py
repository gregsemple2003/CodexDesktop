from __future__ import annotations

import ctypes
from ctypes import wintypes

user32 = ctypes.windll.user32

WM_HOTKEY = 0x0312
PM_REMOVE = 0x0001

MOD_ALT = 0x0001
MOD_CONTROL = 0x0002
MOD_SHIFT = 0x0004
MOD_WIN = 0x0008


class MSG(ctypes.Structure):
    _fields_ = [
        ("hwnd", wintypes.HWND),
        ("message", wintypes.UINT),
        ("wParam", wintypes.WPARAM),
        ("lParam", wintypes.LPARAM),
        ("time", wintypes.DWORD),
        ("pt_x", ctypes.c_long),
        ("pt_y", ctypes.c_long),
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

    def register(self) -> None:
        if self._registered:
            return
        if not user32.RegisterHotKey(
            None,
            self.hotkey_id,
            self.modifiers,
            self.virtual_key,
        ):
            error = ctypes.get_last_error()
            raise OSError(error, f"RegisterHotKey failed for {self.hotkey}")
        self._registered = True

    def poll(self) -> None:
        message = MSG()
        while user32.PeekMessageW(ctypes.byref(message), None, 0, 0, PM_REMOVE):
            if message.message == WM_HOTKEY and int(message.wParam) == self.hotkey_id:
                self.callback()

    def unregister(self) -> None:
        if self._registered:
            user32.UnregisterHotKey(None, self.hotkey_id)
            self._registered = False
