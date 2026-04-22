# 2026-04-21 Human Input Burden And Prime Directive

## Objective

This note answers two questions:

- what the system should optimize for
- how the system should interpret direct human input

## Bottom Line

Direct human input is not normal product traffic. It is failure telemetry.

The safest working rule is:

- treat every direct human input as an error signal
- read that error as system-side insufficiency, not human-side fault
- ask what the system should have inferred, remembered, verified, or continued without needing the human to type

The right emotional frame is also important:

- the human is usually the best sensor of reality on a human-facing task
- frustration is evidence of accumulated exported cost
- sharp correction should be read as high-value truth, not as noise

## Metric

The real thing to minimize is human input burden.

Today, the best durable first approximation is still:

- `InterventionTime`

That remains the most useful measurable proxy for how much extra work the system pushed onto the human.

Your phrasing, `input time / intervention time`, is directionally clear as a forcing idea:

- drive avoidable human attention toward zero

But I do not recommend using that literal ratio as the final scoring rule. Taken literally, it can reward a larger denominator. The safer operational reading is:

- primary objective: minimize direct human input burden
- current proxy: minimize `InterventionTime`
- supporting counters:
  - direct input count
  - repeated correction count
  - wake-up or restart count
  - approval turns
  - restatement turns

If one scalar is needed today, use:

- `HumanInputBurden ~= InterventionTime + weighted_direct_inputs`

and keep `InterventionTime` as the default rough proxy until the rest of that accounting is stable.

## Recovered Local Baseline

The local task history already points in this direction.

- [Task-0007 bootstrap analysis](./2026-04-19-BOOTSTRAP-ANALYSIS.md) says each direct human input is most useful when read as evidence that the system boundary was insufficient.
- [Task-0007 handoff](../../HANDOFF.md) says the standing postmortem question is: `How could the system have inferred the need for this input?`
- [Task-0007 task definition](../../TASK.md) frames the whole task around direct human input as evidence that the system failed to infer something it should have inferred.
- [Jarvis model capture](../Conversations/2026-04-19-JARVIS-MODEL.md) records the original human framing as: treat every input to Codex as a failure signal.

The main refinement from this note is:

- keep the forcing function
- make the compassion and truth rules explicit
- avoid overclaiming clairvoyance

## What The Research Supports

### 1. Interruptions are real cost, not cosmetic friction

Trafton, Altmann, Brock, and Mintz show that interruptions create resumption lag and that preparation and resumption cues matter because returning to the primary task has a real overhead. That maps directly onto “why did the human have to wake the system back up?” and “why is restart supervision expensive?”

Source:

- Trafton et al., *Preparing to resume an interrupted task* (2003): https://gregtrafton.com/papers/preparing.to.resume.pdf

### 2. Mixed-initiative systems must optimize for workload, awareness, and workflow preference, not just raw completion

Gombolay, Bair, Huang, and Shah show that increased autonomy has to be balanced against loss of situational awareness, workload distortion, and disrespect for workflow preferences. That is directly relevant here because a system that “keeps going” while hiding state or ignoring the human’s preferred review path can still raise total human burden.

Source:

- Gombolay et al., *Computational design of mixed-initiative human-robot teaming that considers human factors* (2017): https://doi.org/10.1177/0278364916688255
- Preprint PDF: https://people.csail.mit.edu/gombolay/Publications/Gombolay_IJRR_2017.pdf

### 3. Human intervention cost should be a direct optimization target

HACO is especially relevant because it does not merely accept human intervention as training data. It explicitly treats human cognitive cost and intervention cost as objectives to minimize, and it measures success partly by the decreasing takeover rate over time. That is close to the doctrine needed here.

Source:

- Li, Peng, Zhou, *Efficient Learning of Safe Driving Policy via Human-AI Copilot Optimization* (ICLR 2022): https://arxiv.org/abs/2202.10341

### 4. Each intervention should be propagated forward, not left local to the moment

Predictive Preference Learning from Human Interventions is the cleanest formal statement of the idea that one intervention should teach the system about nearby future states, not just correct the exact moment. The paper also explicitly argues that constant human monitoring creates cognitive burden and that good design should reduce demonstrations needed.

Source:

- Cai, Peng, Zhou, *Predictive Preference Learning from Human Interventions* (NeurIPS 2025): https://arxiv.org/abs/2510.01545

### 5. Tone matters because models mirror tone, and people prefer constructive, well-reasoned responses

Kyrychenko et al. show that users preferred well-reasoned, nuanced responses and that models mirrored linguistic attributes from the user, including toxicity. That matters here for two reasons:

- blame drift is not neutral; it worsens the collaboration loop
- the system should not answer user pain with defensive framing

Source:

- Kyrychenko et al., *Human Preferences for Constructive Interactions in Language Model Alignment* (2025): https://arxiv.org/abs/2503.16480

### 6. Human-AI design guidance already says correction should be cheap and should not compensate for bad default behavior

The survey of industry human-AI guidelines pulls together a consistent set of rules: support efficient correction, act immediately when feedback is received, give people familiar ways to correct errors, and never rely on corrections to make up for low-quality results. That is exactly the right lens for “every input is an error.”

Source:

- Wright et al., *A Comparative Analysis of Industry Human-AI Interaction Guidelines* (2020): https://arxiv.org/abs/2010.11761

## What “Every Input Is An Error” Should Mean

The useful reading is:

- every direct human input is evidence that the system exported too much uncertainty, repair work, or control burden to the human

It should **not** mean:

- the human is wrong for typing
- the system should have been psychic
- every new request is a moral failure

The best truth-seeking reading is:

- default assumption: the system failed
- burden of proof: on the system to show the input was genuine novelty, required approval, or access to hidden external state

That keeps the prime directive sharp without forcing fake certainty.

## Allowed Exceptions

Treat every direct human input as an error signal **until proven otherwise**.

The main exceptions are:

- genuine new objective:
  - the human actually changed the task
- required approval:
  - the system reached a real external gate that should not be auto-crossed
- hidden external state:
  - the needed fact was not available in the repo, packet, or runtime surface
- legitimate preference choice:
  - the human is choosing among several acceptable variants

Even in these cases, the system should still ask:

- could the need for this input have been exposed earlier
- could the choice have been structured better
- could the approval bundle have been made cheaper to review

## Derived Design Rules

### 1. Do not blame the human for boundary repair

Bad frame:

- the human added constraints
- the human changed the plan
- the human moved the goalposts

Better frame:

- the system failed to ground a latent boundary
- the system forgot or violated a known boundary
- the human had to restore reality to the loop

### 2. Restart supervision is severe failure

Messages like:

- `continue now`
- `why are you messaging me`
- `do not hand this back to me`

should be treated as high-severity evidence that the system exported scheduler work to the human.

### 3. False proof of done is severe failure

If the system claims closure on the wrong lane, with the wrong evidence, or in the wrong answer shape, it has already created rework. A later correction is not minor cleanup. It is proof that the default closure surface was wrong.

### 4. Repeated corrections must become durable changes

A correction that recurs is not “context.” It is a failed learning loop. The system should promote it into:

- memory
- a verifier
- a gate
- a task
- a packet rule

### 5. Frustration is signal

Do not normalize away sharp language. It often means the human has already paid the cost several times.

## Recommended Prime Directive Wording

Use this wording when the system needs the shortest durable version:

> Prime directive: minimize human input burden. Every direct human input is an error signal until the system proves it was genuine novelty, required approval, or hidden external state. Default assumption: the system failed to infer, retain, verify, continue, or present something it should have handled without the human typing. Do not blame the human for boundary repair. Ask what the system should have known earlier and what durable change would make this input unnecessary next time.

## Recommended Postmortem Questions

After each direct human input, ask:

1. What work did this input force the human to do?
2. What should the system have known earlier?
3. Was the failure about inference, memory, proof, continuation, answer shape, or scope?
4. What durable change would make this input unnecessary next time?
5. If this was true novelty, what evidence proves that?

## Best Current Recommendation

The best operational doctrine is:

- `InterventionTime` remains the first approximation to human cost
- every direct human input is an error signal
- default to system fault, not human fault
- require proof before classifying an input as genuine novelty
- measure improvement by whether repeated intervention goes down over time

That is the cleanest way to make the system accept its new role without becoming dishonest about what it could have known.

## Sources

Local sources:

- [Bootstrap analysis](./2026-04-19-BOOTSTRAP-ANALYSIS.md)
- [Research analysis root](../../RESEARCH-ANALYSIS.md)
- [Task handoff](../../HANDOFF.md)
- [Task definition](../../TASK.md)
- [Jarvis model capture](../Conversations/2026-04-19-JARVIS-MODEL.md)
- [InterventionTime prioritization](../../InterventionTime/PRIORITIZATION.md)
- [Topic 08 autonomy and safety policy](./TOPIC-08-AUTONOMY-AND-SAFETY-POLICY.md)
- [Topic 09 evaluation loops](./TOPIC-09-EVALUATION-LOOPS.md)
- [Predictive Preference Learning local sidecar](../Sources/Arxiv/2025-10-02-PREDICTIVE-PREFERENCE-LEARNING-HUMAN-INTERVENTIONS.md)

Primary external sources:

- Trafton et al. (2003), *Preparing to resume an interrupted task*: https://gregtrafton.com/papers/preparing.to.resume.pdf
- Gombolay et al. (2017), *Computational design of mixed-initiative human-robot teaming that considers human factors*: https://doi.org/10.1177/0278364916688255
- Li et al. (2022), *Efficient Learning of Safe Driving Policy via Human-AI Copilot Optimization*: https://arxiv.org/abs/2202.10341
- Wright et al. (2020), *A Comparative Analysis of Industry Human-AI Interaction Guidelines*: https://arxiv.org/abs/2010.11761
- Kyrychenko et al. (2025), *Human Preferences for Constructive Interactions in Language Model Alignment*: https://arxiv.org/abs/2503.16480
- Cai et al. (2025), *Predictive Preference Learning from Human Interventions*: https://arxiv.org/abs/2510.01545
