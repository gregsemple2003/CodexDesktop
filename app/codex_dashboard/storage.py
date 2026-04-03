from __future__ import annotations

import sqlite3
from datetime import datetime
from pathlib import Path

from .models import TokenEvent


def connect(db_path: Path) -> sqlite3.Connection:
    db_path.parent.mkdir(parents=True, exist_ok=True)
    connection = sqlite3.connect(db_path)
    connection.row_factory = sqlite3.Row
    return connection


def initialize_db(connection: sqlite3.Connection) -> None:
    connection.executescript(
        """
        PRAGMA journal_mode=WAL;

        CREATE TABLE IF NOT EXISTS file_cursors (
            path TEXT PRIMARY KEY,
            last_offset INTEGER NOT NULL,
            last_size INTEGER NOT NULL,
            updated_at TEXT NOT NULL
        );

        CREATE TABLE IF NOT EXISTS token_events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            session_path TEXT NOT NULL,
            line_offset INTEGER NOT NULL,
            event_timestamp TEXT NOT NULL,
            total_tokens INTEGER NOT NULL,
            input_tokens INTEGER NOT NULL,
            cached_input_tokens INTEGER NOT NULL,
            output_tokens INTEGER NOT NULL,
            reasoning_output_tokens INTEGER NOT NULL,
            cumulative_total_tokens INTEGER NOT NULL,
            weekly_used_percent REAL,
            weekly_window_minutes INTEGER,
            weekly_resets_at INTEGER,
            raw_json TEXT NOT NULL,
            UNIQUE(session_path, line_offset)
        );
        """
    )
    connection.commit()


def load_cursor(connection: sqlite3.Connection, path: str) -> tuple[int, int]:
    row = connection.execute(
        "SELECT last_offset, last_size FROM file_cursors WHERE path = ?",
        (path,),
    ).fetchone()
    if row is None:
        return 0, 0
    return int(row["last_offset"]), int(row["last_size"])


def save_cursor(
    connection: sqlite3.Connection,
    path: str,
    last_offset: int,
    last_size: int,
) -> None:
    connection.execute(
        """
        INSERT INTO file_cursors(path, last_offset, last_size, updated_at)
        VALUES (?, ?, ?, ?)
        ON CONFLICT(path) DO UPDATE SET
            last_offset = excluded.last_offset,
            last_size = excluded.last_size,
            updated_at = excluded.updated_at
        """,
        (path, last_offset, last_size, datetime.now().isoformat()),
    )


def insert_event(connection: sqlite3.Connection, event: TokenEvent) -> bool:
    cursor = connection.execute(
        """
        INSERT OR IGNORE INTO token_events(
            session_path,
            line_offset,
            event_timestamp,
            total_tokens,
            input_tokens,
            cached_input_tokens,
            output_tokens,
            reasoning_output_tokens,
            cumulative_total_tokens,
            weekly_used_percent,
            weekly_window_minutes,
            weekly_resets_at,
            raw_json
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        """,
        (
            event.session_path,
            event.line_offset,
            event.event_timestamp.isoformat(),
            event.total_tokens,
            event.input_tokens,
            event.cached_input_tokens,
            event.output_tokens,
            event.reasoning_output_tokens,
            event.cumulative_total_tokens,
            event.weekly_used_percent,
            event.weekly_window_minutes,
            event.weekly_resets_at,
            event.raw_json,
        ),
    )
    return cursor.rowcount > 0


def load_events_since(
    connection: sqlite3.Connection,
    since: datetime,
) -> list[TokenEvent]:
    rows = connection.execute(
        """
        SELECT *
        FROM token_events
        WHERE event_timestamp >= ?
        ORDER BY event_timestamp ASC
        """,
        (since.isoformat(),),
    ).fetchall()
    events: list[TokenEvent] = []
    for row in rows:
        events.append(
            TokenEvent(
                session_path=str(row["session_path"]),
                line_offset=int(row["line_offset"]),
                event_timestamp=datetime.fromisoformat(str(row["event_timestamp"])),
                total_tokens=int(row["total_tokens"]),
                input_tokens=int(row["input_tokens"]),
                cached_input_tokens=int(row["cached_input_tokens"]),
                output_tokens=int(row["output_tokens"]),
                reasoning_output_tokens=int(row["reasoning_output_tokens"]),
                cumulative_total_tokens=int(row["cumulative_total_tokens"]),
                weekly_used_percent=(
                    float(row["weekly_used_percent"])
                    if row["weekly_used_percent"] is not None
                    else None
                ),
                weekly_window_minutes=(
                    int(row["weekly_window_minutes"])
                    if row["weekly_window_minutes"] is not None
                    else None
                ),
                weekly_resets_at=(
                    int(row["weekly_resets_at"])
                    if row["weekly_resets_at"] is not None
                    else None
                ),
                raw_json=str(row["raw_json"]),
            )
        )
    return events


def count_events(connection: sqlite3.Connection) -> int:
    row = connection.execute("SELECT COUNT(*) AS total FROM token_events").fetchone()
    return int(row["total"])
