# Prompt Input

This file captures the exact batch-specific input text used to generate and later modify the `HumanInterventionTime` packet in this directory.

Recovered on April 20, 2026 from Codex session transcripts via `codex-session-search`.

Important truth:
the current packet is not purely worker-generated. The present outputs reflect an initial worker run, a rewrite worker run, and later parent-thread shell-command post-processing.

## Initial Worker Launch Input

Transcript source:
[rollout-2026-04-20T12-05-16-019daba3-8fdb-7a82-874a-22ae0e1fe882.jsonl#L6](/c:/Users/gregs/.codex/sessions/2026/04/20/rollout-2026-04-20T12-05-16-019daba3-8fdb-7a82-874a-22ae0e1fe882.jsonl#L6)

Historical note:
the quoted launch and post-processing text below predates promotion into the shared `Reports\Interventions` root, so the task-local generation path is preserved verbatim.

```text
You are responsible only for files under C:\Agent\CodexDashboard\Tracking\Task-0007\InterventionTime\2026-04-19-thirdperson-v0\. Other agents may be working elsewhere in the repo; do not revert or touch their files. Read C:\Agent\CodexDashboard\Tracking\Task-0007\InterventionTime\2026-04-19-thirdperson-v0\RUN-MEASURE.md and follow it exactly. Stay context-free: do not use memory, do not search outside the explicitly allowed files, and write only the required outputs into that folder. When finished, report the exact files you changed and the key totals.
```

## Rewrite Worker Launch Input

Transcript source:
[rollout-2026-04-20T12-12-23-019dabaa-147f-7c63-b7c4-600190bc73c6.jsonl#L6](/c:/Users/gregs/.codex/sessions/2026/04/20/rollout-2026-04-20T12-12-23-019dabaa-147f-7c63-b7c4-600190bc73c6.jsonl#L6)

```text
You are responsible only for files under C:\Agent\CodexDashboard\Tracking\Task-0007\InterventionTime\2026-04-19-thirdperson-v0\. Other agents may be working elsewhere in the repo; do not revert or touch their files. Re-read C:\Agent\CodexDashboard\Tracking\Task-0007\InterventionTime\2026-04-19-thirdperson-v0\RUN-MEASURE.md and follow it exactly. Rewrite the required output files in that folder using the tightened explicit-wait rule. Stay context-free: do not use memory, do not search outside the explicitly allowed files, and write only the required outputs into that folder. When finished, report the exact files you changed and the key totals.
```

## Parent-Thread Post-Processing: Explicit Wait / Approval Gate Patch

Transcript source:
[rollout-2026-04-19T13-44-02-019da6d7-9f4b-73b3-8f9b-bc4c7dcf4a4e.jsonl#L2288](/c:/Users/gregs/.codex/sessions/2026/04/19/rollout-2026-04-19T13-44-02-019da6d7-9f4b-73b3-8f9b-bc4c7dcf4a4e.jsonl#L2288)

```powershell
$dir='C:\Agent\CodexDashboard\Tracking\Task-0007\InterventionTime\2026-04-19-thirdperson-v0';
$eventsPath=Join-Path $dir 'EVENT-MEASUREMENTS.jsonl';
$stallPath=Join-Path $dir 'STALL-EVENTS.json';
$summaryPath=Join-Path $dir 'SUMMARY.json';
$events = Get-Content -Path $eventsPath | ForEach-Object { $_ | ConvertFrom-Json };
$target = $events | Where-Object { $_.event_id -eq 'hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919' };
if (-not $target) { throw 'Target event not found'; }
$target.intervention_kind = 'answer_to_question';
$target.prior_ai_state.explicit_wait_state = $true;
$target.stall_basis = 'explicit_wait_state';
$target.stall_loss_seconds = 0.0;
$target.intervention_time_seconds = [math]::Round([double]$target.typing_seconds_estimate, 3);
$target.notes = @('manual_chars from raw_text', 'explicit approval gate; stall loss suppressed under tightened rule');
$events | ForEach-Object { $_ | ConvertTo-Json -Compress -Depth 8 } | Set-Content -Path $eventsPath;
$stallEvents = @($events | Where-Object { [double]$_.stall_loss_seconds -gt 0 });
$stallPayload = [ordered]@{
  schema_version = 'intervention_time.stalls.v0'
  packet_label = $events[0].packet_label
  stall_event_count = $stallEvents.Count
  stall_events = @($stallEvents | ForEach-Object {
    [ordered]@{
      event_id = $_.event_id
      captured_at_local = $_.captured_at_local
      intervention_kind = $_.intervention_kind
      idle_gap_seconds = $_.prior_ai_state.idle_gap_seconds
      stall_grace_seconds = $_.prior_ai_state.stall_grace_seconds
      stall_loss_seconds = $_.stall_loss_seconds
      evidence = @($_.evidence | ForEach-Object { [ordered]@{ ref = $_.ref; excerpt = $_.excerpt } })
    }
  })
};
$stallPayload | ConvertTo-Json -Depth 8 | Set-Content -Path $stallPath;
$counts = [ordered]@{ new_request = 0; correction = 0; boundary_reset = 0; answer_to_question = 0; wake_up = 0; other = 0 };
foreach ($event in $events) { if ($counts.Contains($event.intervention_kind)) { $counts[$event.intervention_kind] = [int]$counts[$event.intervention_kind] + 1 } }
$totalManualChars = [int](($events | Measure-Object -Property manual_chars -Sum).Sum);
$totalTypingSeconds = [math]::Round([double](($events | Measure-Object -Property typing_seconds_estimate -Sum).Sum), 3);
$totalStallSeconds = [math]::Round([double](($events | Measure-Object -Property stall_loss_seconds -Sum).Sum), 3);
$totalInterventionSeconds = [math]::Round([double](($events | Measure-Object -Property intervention_time_seconds -Sum).Sum), 3);
$topStalls = @($stallEvents | Sort-Object -Property stall_loss_seconds -Descending | Select-Object -First 5 | ForEach-Object {
  $preview = ($_.evidence | Where-Object { $_.source -eq 'source_packet' } | Select-Object -First 1).excerpt;
  [ordered]@{
    event_id = $_.event_id
    stall_loss_seconds = $_.stall_loss_seconds
    captured_at_local = $_.captured_at_local
    preview = $preview
  }
});
$summaryPayload = [ordered]@{
  schema_version = 'intervention_time.summary.v0'
  packet_label = $events[0].packet_label
  repo_ref = 'C:\Agent\ThirdPerson'
  local_date = '2026-04-19'
  event_count = $events.Count
  stall_event_count = $stallEvents.Count
  typing_rate_chars_per_second = $events[0].typing_rate_chars_per_second
  stall_grace_seconds = $events[0].prior_ai_state.stall_grace_seconds
  totals = [ordered]@{
    manual_chars = $totalManualChars
    typing_seconds_estimate = $totalTypingSeconds
    stall_loss_seconds = $totalStallSeconds
    intervention_time_seconds = $totalInterventionSeconds
  }
  counts_by_intervention_kind = $counts
  top_stall_events = $topStalls
  notes = @(
    'Manual characters are counted from parsed_envelope.request_text when present; otherwise raw_text.',
    'Stall loss is charged only on wake_up events after the 60-second grace window.',
    'Event hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702 was reclassified to boundary_reset and set to zero stall loss because it followed an explicit wait gate.',
    'Event hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919 was reclassified to answer_to_question and set to zero stall loss because it followed an approval gate.'
  )
};
$summaryPayload | ConvertTo-Json -Depth 8 | Set-Content -Path $summaryPath;
Write-Output 'patched intervention-time outputs'
```

## Parent-Thread Post-Processing: Typing-Rate Recompute

Transcript source:
[rollout-2026-04-19T13-44-02-019da6d7-9f4b-73b3-8f9b-bc4c7dcf4a4e.jsonl#L2373](/c:/Users/gregs/.codex/sessions/2026/04/19/rollout-2026-04-19T13-44-02-019da6d7-9f4b-73b3-8f9b-bc4c7dcf4a4e.jsonl#L2373)

```powershell
$dir='C:\Agent\CodexDashboard\Tracking\Task-0007\InterventionTime\2026-04-19-thirdperson-v0';
$eventsPath=Join-Path $dir 'EVENT-MEASUREMENTS.jsonl';
$stallPath=Join-Path $dir 'STALL-EVENTS.json';
$summaryPath=Join-Path $dir 'SUMMARY.json';
$typingRate = 1.0;
$events = Get-Content -Path $eventsPath | ForEach-Object { $_ | ConvertFrom-Json };
foreach ($event in $events) {
  $event.typing_rate_chars_per_second = $typingRate;
  $event.typing_seconds_estimate = [math]::Round(([double]$event.manual_chars / $typingRate), 3);
  $event.intervention_time_seconds = [math]::Round(([double]$event.typing_seconds_estimate + [double]$event.stall_loss_seconds), 3);
}
$events | ForEach-Object { $_ | ConvertTo-Json -Compress -Depth 8 } | Set-Content -Path $eventsPath;
$stallPayload = Get-Content -Raw -Path $stallPath | ConvertFrom-Json;
$summaryPayload = Get-Content -Raw -Path $summaryPath | ConvertFrom-Json;
$totalManualChars = [int](($events | Measure-Object -Property manual_chars -Sum).Sum);
$totalTypingSeconds = [math]::Round([double](($events | Measure-Object -Property typing_seconds_estimate -Sum).Sum), 3);
$totalStallSeconds = [math]::Round([double](($events | Measure-Object -Property stall_loss_seconds -Sum).Sum), 3);
$totalInterventionSeconds = [math]::Round([double](($events | Measure-Object -Property intervention_time_seconds -Sum).Sum), 3);
$summaryPayload.typing_rate_chars_per_second = $typingRate;
$summaryPayload.totals.manual_chars = $totalManualChars;
$summaryPayload.totals.typing_seconds_estimate = $totalTypingSeconds;
$summaryPayload.totals.stall_loss_seconds = $totalStallSeconds;
$summaryPayload.totals.intervention_time_seconds = $totalInterventionSeconds;
$summaryPayload.notes = @(
  'Manual characters are counted from parsed_envelope.request_text when present; otherwise raw_text.',
  'Typing rate is intentionally conservative at 1.0 chars/second so the baseline partially absorbs human reading and thinking time.',
  'Stall loss is charged only on wake_up events after the 60-second grace window.',
  'Event hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702 was reclassified to boundary_reset and set to zero stall loss because it followed an explicit wait gate.',
  'Event hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919 was reclassified to answer_to_question and set to zero stall loss because it followed an approval gate.'
);
$summaryPayload | ConvertTo-Json -Depth 8 | Set-Content -Path $summaryPath;
Write-Output 'recomputed intervention-time packet for typing_rate=1.0'
```

## Canonical Recipe And Local Context

- [PROMPT.md](../../../../../../../Prompts/Interventions/HumanInterventionTime/PROMPT.md)
- [WORKFLOW.md](../../../../../../../Prompts/Interventions/HumanInterventionTime/WORKFLOW.md)
- [DISPATCH-PATTERN.md](../../../../../../../Prompts/Interventions/HumanInterventionTime/DISPATCH-PATTERN.md)
- [LOCAL-CONTEXT.json](./LOCAL-CONTEXT.json)
