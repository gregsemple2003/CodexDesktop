# 2025-10-02-predictive-preference-learning-human-interventions.pdf

Source PDF: `2025-10-02-predictive-preference-learning-human-interventions.pdf`

Page count: 21

## Page 1

Predictive Preference Learning from Human
Interventions
Haoyuan Cai, Zhenghao Peng, Bolei Zhou
Department of Computer Science,
University of California, Los Angeles
Abstract
Learning from human involvement aims to incorporate the human subject to moni-
tor and correct agent behavior errors. Although most interactive imitation learning
methods focus on correcting the agent’s action at the current state, they do not adjust
its actions in future states, which may be potentially more hazardous. To address
this, we introduce Predictive Preference Learning from Human Interventions (PPL),
which leverages the implicit preference signals contained in human interventions
to inform predictions of future rollouts. The key idea of PPL is to bootstrap each
human intervention into L future time steps, called the preference horizon, with the
assumption that the agent follows the same action and the human makes the same
intervention in the preference horizon. By applying preference optimization on
these future states, expert corrections are propagated into the safety-critical regions
where the agent is expected to explore, significantly improving learning efficiency
and reducing human demonstrations needed. We evaluate our approach with ex-
periments on both autonomous driving and robotic manipulation benchmarks and
demonstrate its efficiency and generality. Our theoretical analysis further shows
that selecting an appropriate preference horizon L balances coverage of risky states
with label correctness, thereby bounding the algorithmic optimality gap. Demo and
code are available at:https://metadriverse.github.io/ppl.
1 Introduction
Effectively leveraging human demonstrations to teach and align autonomous agents remains a central
challenge in both Reinforcement Learning (RL) [46] and Imitation Learning (IL) [17]. In the literature
of RL and more recent RL from Human Feedback (RLHF), the agent explores the environment
through trial and error or under human feedback guidance, and the learning process hinges on a
carefully crafted reward function that reflects human preferences. However, RL algorithms often
require a large number of environment interactions to learn stable policies, and their exploration
can lead agents to dangerous or task-irrelevant states [40, 27]. In contrast, IL methods train agents
to emulate human behavior using offline demonstrations from experts. Nevertheless, IL agents
are susceptible to distributional shift because the offline dataset may lack corrective samples in
safety-critical or out-of-distribution states [35, 32, 3, 47].
Interactive Imitation Learning (IIL) [2, 34, 15, 42, 28, 41, 19, 20] incorporates human participants to
intervene in the training process and provide online demonstrations. Such methods have improved
alignment and learning efficiency in a wide variety of tasks, including robot manipulation [ 7, 8],
autonomous driving [27, 28], and even the strategy game StarCraft II [ 39]. One line of research
on confidence-based IIL designs various task-specific criteria to request human help, including
uncertainty estimation [23] and confidence in completing the task [4, 38]. In contrast, an increasing
body of work focuses on learning from active human involvement, where human subjects actively
intervene and provide demonstrations during training when the agent makes mistakes [15, 42, 22, 17,
28]. Compared to confidence-based IIL, active human involvement can ensure training safety [27]
39th Conference on Neural Information Processing Systems (NeurIPS 2025).
arXiv:2510.01545v2  [cs.LG]  15 Oct 2025

## Page 2

You will crash soon! Let me help you.
1) Trajectory Prediction Facilitating Human Intervention
Can I turn right here?
2) Learning Human Preferences in Forecasted Rollouts
Human rejects my right-turn.
I understand human preferences in these states. 
Figure 1: Our Predictive Preference Learning from Human Interventions. (Top) Our approach
forecasts the agent’s upcoming trajectory (the red dotted path) and visualizes it for the human expert,
who will intervene if the forecasted path indicates an upcoming failure. (Bottom) A single intervention
is then interpreted as hypothesized preference signals across the predicted states. These signals reflect
the agent’s imputed imagination of what the expert would prefer, guiding the policy to avoid the risky
maneuver in similar future contexts. This integration of proactive forecasting and preference learning
accelerates policy improvement and reduces the total number of expert interventions required.
and does not require carefully designing human intervention criteria for each task [ 12]. However,
these methods require the human expert to monitor the entire training process, predict the agent’s
future trajectories, and intervene immediately in safety-critical states [ 28], imposing a significant
cognitive burden on the human participant. In addition, these methods often correct the agent’s
behavior only at the current intervened state, penalizing undesired actions step by step. For instance,
in HG-DAgger [15], the agent is optimized to mimic human actions solely at the states where
interventions occur. In practice, it is intuitive that the agent may repeat similar mistakes in the
consecutive future steps t+ 1,· · ·, t+L following an error at step t. As a result, the expert must
repeatedly provide corrective demonstrations in these regions, compromising training efficiency [17].
In this work, we propose a novel Interactive Imitation Learning algorithm,Predictive Preference
Learningfrom Human Interventions ( PPL), to learn from active human involvement. As shown
in Fig. 1, our approach has two key designs: First, we employ an efficient rollout-based trajectory
prediction model to forecast the agent’s future states. These predicted rollouts are visualized in real
time for the user, helping human supervisors proactively determine when an intervention is necessary.
Second, our algorithm leverages preference learning on the predicted future trajectories to further
improve the sample efficiency and reduce the expert demonstrations needed. Such designs bring three
strengths: (1) They mitigate the distributional shift problem in IIL and improve training efficiency. By
incorporating anticipated future states into the training process, our method constructs a richer dataset,
especially in safety-critical situations. This expanded dataset offers more information than expert
corrective demonstrations in human-intervened states only. (2) The preference learning reduces the
agent’s visits to dangerous states, thus suppressing human interventions in safety-critical situations.
(3) By visualizing the agent’s predicted future trajectories in the user interface, we significantly
reduce the cognitive burden on the human supervisor to constantly anticipate the agent’s behavior.
Our contributions can be summarized as follows:
1. We introduce a novel Interactive Imitation Learning (IIL) algorithm that leverages trajectory
prediction to inform human intervention and employs preference learning to deter the agent
from returning to dangerous states.
2. We evaluate our algorithm on the MetaDrive [16] and Robosuite [49] benchmarks, using
both neural experts and real human participants, showing that PPL requires fewer expert
monitoring efforts and demonstrations to achieve near-optimal policies.
3. We present a theoretical analysis that derives an upper bound on the performance gap
of our approach. This bound highlights that the efficacy of our method lies in reducing
distributional shifts while preserving the quality of preference data.
2

## Page 3

2 Related Work
Learning from Human Involvement.Many works incorporate human involvement in the training
loop to provide corrective actions in dangerous or repetitive states. For example, Human-Gated
DAgger (HG-DAgger) [15], Ensemble-DAgger [23], Thrifty-DAgger [12], Sirius [19], and Inter-
vention Weighted Regression (IWR) [22] perform imitation learning on human intervention data.
These methods do not leverage data collected by agents or suppress undesired actions likely to be
intervened by humans, leading to the agent’s susceptibility to entering hazardous states and thus
harming sample efficiency. EGPO [27], PVP [28], and AIM [2] design proxy cost or value functions
to suppress the frequency of human involvement. However, these approaches still require human
supervisors to continuously monitor the agent’s behavior throughout training and anticipate potential
failures that may necessitate intervention. This continuous oversight imposes a significant cognitive
load on the human expert and can limit scalability. Furthermore, these methods do not exploit the
agent’s predicted future trajectories that the expert might identify as potentially leading to undesirable
outcomes, which necessitates repeated corrective demonstrations in such situations.
Preference-Based RL.A large body of work focuses on learning human preferences by ranking
pairs of trajectories generated by the agent [6, 9, 34, 44, 37, 26]. One prominent paradigm, reinforce-
ment learning from human feedback (RLHF), first trains a reward model on offline human preference
data and then uses that model to guide policy optimization [6, 25, 43]. RLHF has achieved impres-
sive results in domains ranging from Atari games [6] to large language models [25]. Alternatively,
methods such as Direct Preference Optimization (DPO) [31], Contrastive Preference Optimization
(CPO) [45], and related variants [1, 24] bypass explicit reward-model training and instead directly
optimize the policy to satisfy preference labels via a classification loss.
However, applying RLHF and DPO to real-time control problems faces challenges due to the need for
extensive human labeling of preference data [31]. These labels are inherently subjective and prone
to noise [37]. Moreover, acquiring a high-quality preference dataset and achieving near-optimal
policies often requires a substantial number of environment samples, thereby imposing a considerable
burden on human experts [10]. In contrast, our framework elicits preferences in an online, interactive
manner: experts review the agent’s predicted future trajectory at each decision point and intervene
when a failure is anticipated; these interventions are then converted into contrastive preference labels.
This real-time preference collection enables the policy to adapt continuously to the evolving state
distribution and to receive targeted feedback precisely where it is most needed. In summary, our
approach PPL bridges preference-based RL and imitation learning by demonstrating that DPO-style
alignment techniques can be effectively adapted to control problems within an interactive imitation
learning framework.
3 Problem Formulation
In this section, we introduce our settings of interactive imitation learning environments. We use
the Markov decision process (MDP) M=⟨S,A,P, r, γ, d 0⟩ to model the environment, which
contains a state space S, an action space A, a state transition function P:S × A → S , a reward
function r:S × A →[R min, Rmax], a discount factor γ∈(0,1) , and an initial state distribution
d0 :S →[0,1] . We denote π(a|s) :S × A →[0,1] as a stochastic policy. Reinforcement
learning (RL) aims to learn anovice policy πn(a|s) that maximizes the expected cumulative return
J(π n) = E
τ∼P πn
[
∞P
t=0
γtr(st, at)], wherein τ= (s 0, a0, s1, a1, ...) is the trajectory sampled from
trajectory distribution Pπn induced by πn, d0 and P. We also define the discounted state distribution
under policy πn as dπn (s) = (1−γ) E
τ∼P πn
[
∞P
t=0
γtI[st =s]]. In this work, we consider the reward-free
setting where the agent has no access to the task reward functionr(s, a).
In imitation learning (IL), we assume that the human expert behavior ah follows ahuman policy
πh(a|s) . The agent aims to learn πn from human expert trajectories τh ∼P πh, and it needs
to optimize πn to close the gap between τn ∼P πn and τh. Prior works on imitation learning
have shown that using an offline expert demonstration dataset may lead to poor performance due
to out-of-distribution states [ 36, 34, 41]. Therefore, interactive imitation learning (IIL) methods
incorporate a human expert into the training loop to provide online corrective demonstrations, making
3

## Page 4

the state distribution of expert data more similar to that of the novice policy [28, 35]. During training,
the human expert monitors the agent and can intervene and take control if the agent’s actionan at
the current state s violates the human’s desired behavior or leads to a dangerous situation. We use
the deterministic intervention policy I(s, an) :S × A → {0,1} to model the human’s intervention
behavior, where the agent’s action follows the novice policyan ∼π n(· |s) , and the human subject
takes control whenI(s, a n) = 1.
With the notations above, the agent’s actual trajectories during training are derived from the following
sharedbehavior policy
πb(a|s) =π n(a|s)(1−I(s, a)) +π h(a|s)G(s),(1)
wherein G(s) =
R
a′∈A I(s, a′)πn(a′ |s)da ′ is the probability of the agent taking an action an that
will be rejected and intervened by the human expert.
Preference Alignment.Recent works on preference-based RL have also leveraged offline preference
datasets to learn human-aligned policies [31, 45, 1]. Given an offline preference dataset Dpref where
each preference data (s, a+, a−)∈ D pref means that the human expert prefers the action a+ over
a− at state s, we can learn an agent policy πn that aligns with the human preference model. The
Contrastive Preference Optimization method [45] uses the following objective to train an agent policy
πθ from the preference datasetD pref:
Lpref(πθ) =− E
(s,a+,a−)∼Dpref

logσ
 
βlogπ θ(a+ |s)−βlogπ θ(a− |s)

,(2)
whereσ(·)is the Sigmoid function, andβ >0is a hyperparameter.
Trajectory Prediction Model.In this work, we allow the agent to access a short-term trajectory
prediction model f(s, a n, H). Given the current state s and the agent’s actionan, we can predict the
agent’s trajectoryf(s, a n, H) = (s,˜s1,· · ·,˜s H ) in the next H steps, where ˜si the predicted state that
the agent will reach if the agent applies the action an for i steps from the state s. The implementation
detail offis in Sec. 4.3.
4 Method
4.1 Predictive Preference Learning from Human Interventions (PPL)
We propose PPL (Fig. 2), an efficient interactive imitation learning method that emulates the human
policy with fewer expert demonstrations and less cognitive effort. The key idea of PPL is to learn
human preferences from data generated by a future-trajectory prediction model. We illustrate the
human-agent interactions in Fig. 2 (left) and how PPL infers human preference in Fig. 2 (right).
During training, the human subject monitors the agent-environment interaction in each state s (Fig. 2
(left)). The novice policy πn suggests an action an for the current state s. Instead of executing
an immediately, we query the trajectory prediction model f(s, a n, H) to obtain a predicted rollout
τ=f(s, a n, H) = (s,˜s1,· · ·,˜s H ), which we visualize for the human expert. The expert then uses
τ to determine whether the agent will fail in the next H steps, such as crashing into vehicles or
going off the road. If so, the expert will provide corrective actions ah ∼π h(s) for the next H steps,
depicted by the blue trajectory in Fig. 2. If the expert believes no intervention is needed, the agent
continues to use its own policyπ n for the nextHsteps.
We introduce preference learning on the predicted trajectories because it is difficult to learn corrective
behavior purely from the expert’s demonstrations in safe states. By visualizing predicted rollouts,
experts can anticipate unsafe trajectories before the agent actually enters them and intervene preemp-
tively. As a result, the state distribution covered by these early interventions differs substantially from
the on-policy distribution of the novice policy, creating a distributional shift that standard imitation or
on-policy correction cannot address. Therefore, instead of relying solely on expert demonstrations,
we collect preference labels over the predicted rollouts (Fig. 2 Right) so that the agent can learn the
correct behavior in those risky states.
Whenever the expert intervenes at state s, we interpret this as indicating that continuing with an
would lead to unsafe or undesirable outcomes along the predicted trajectory. As shown in Fig. 2
(right), to capture this preference, we assume the expert prefers ah over an at state s and each of
4

## Page 5

Preference Buffer
You should do  in state , not .
Human prefers  over  at .
Human prefers  over  at .
1) Human intervenes to prevent a future crash.
2) Agent infers human preference in predicted future states.Human DemonstrationAgent's Exploratory 
Trajectory Agent's Predicted 
Trajectory
 Human Buffer 
Figure 2: Illustration of Predictive Preference Learning. (Left) At each decision point, the agent
proposes an action, and its future trajectory is predicted and visualized. The human expert reviews
this rollout and intervenes only when a potential failure is anticipated. The intervention is recorded
alongside the state into the human buffer Dh for behavioral cloning. (Right) Each recorded interven-
tion is then converted into contrastive preference pairs over the predicted future states ˜s1,· · ·,˜s L.
These preference tuples are stored in a preference buffer Dpref and used to train the policy via a
contrastive classification loss, propagating expert intents into regions the agent is likely to explore.
the first L predicted states ˜s1,· · ·,˜s L for some preference horizon L≤H . For each i≤L , we
add the tuple (˜si, a+ =a h, a− =a n) to the preference dataset Dpref. We note that in each tuple
(˜si, a+ =a h, a− =a n), both ah and an are sampled at the current state s, not the predicted future
states ˜si, because the exact human corrective actions at hypothetical future states are not directly
observable. Still, the expert intervention at state s implies that applying ah at the predicted states
˜s1, . . . ,˜sL, rather than continuing with an, is more likely to prevent the dangerous outcome in the
end of the predicted trajectory (˜sH). Hence, our construction of the preference dataset ensures that it
faithfully captures the expert’s corrective intent across the predicted horizon.
The preference horizon L controls the length over which we elicit preferences in the predicted
trajectory. A small L may fail to capture enough risky states, while a large L risks applying
preferences where the corrective action ah at state s no longer matches what an expert would do in
those imagined states ˜si. In Theorem 4.1, we prove that under mild assumptions, the performance
gap of our learned policy is bounded by terms reflecting the state distribution shift and the quality
of the preference labels, implying that an ideal preference horizonLshould balance these two error
terms. We also illustrate how the choice ofLaffects the performance of PPL in Fig. 8.
We train the novice policy πn using two complementary objectives. First, we apply a behavioral
cloning loss on expert demonstrationsD h:
LBC(πθ) =− E
(s,ah)∼Dh
[logπ θ(ah |s)].(3)
Second, inspired by Contrastive Preference Optimization (CPO) [ 45], we use the preference-
classification loss Eq. 2 over the predicted states in Dpref. The final loss of the agent policy πθ
is evaluated as
L(πθ) =L pref(πθ) +L BC(πθ)
=− E
(s,a+,a−)∼Dpref

logσ
 
βlogπ θ(a+ |s)−βlogπ θ(a− |s)

− E
(s,ah)∼Dh
[logπ θ(ah |s)].
(4)
The workflow of our method PPL is summarized in Alg. 1.
4.2 Analysis
We prove that the performance gap between the human policy πh and the agent policy πn can be
bounded by the following three error terms: 1) the state distribution shift δdist, 2) the quality of the
preference labelsδ pref, and 3) the optimization errorϵ.
The first error term is defined as δdist =D TV(dπn , dpref), where dπn (s) is the discounted state
distribution of the agent’s policy πn, and dpref(s) =|D pref|−1
E(s′,a+,a−)∼Dpref I[s′ =s] . Here,
5

## Page 6

Test Success Rate of MetaDrive
PPL (Ours) PVP 
2k4k6k8kSteps00.20.40.60.8Success Rate
Test Success Rate of Table Wiping
PPL (Ours) PVP 
1k2k3k4kSteps00.20.40.60.8Success Rate
Test Success Rate of Nut Assembly
PPL (Ours) PVP 
50k100k150k200kSteps0.10.20.30.40.5Success Rate
Figure 3: The test-time performance curve of PPL and the IIL counterpart PVP [ 28] under three
different environments. The x-coordinate is the number of environment interactions, and the y-
coordinate is the agent’s success rate in a held-out test environment, where the evaluation is conducted
without expert involvement. Compared to the IIL counterpart, our approach achieves much higher
learning efficiency and reduces the expert’s efforts needed.
DTV(P, Q) = 1
2 |P−Q| 1 is the total variation distance between two distributions. This error term
quantitatively measures the difference between the states actually visited by the agent and those
contained in the preference dataset.
The second error term is defined as δpref = E
s∼dpref
DTV(ρs
ideal, ρs
pref), which arises from the mis-
alignment of the positive actions in the preference dataset, as the human action ah in each tuple
(˜si, ah, an)∈ D pref is sampled in state s instead of state ˜si. That is, this error reflects the assumption
that the human would still apply the same corrective actionah in a hypothetical future state ˜si reached
after executing an for i steps, which may not perfectly match what the expert would actually do. For
any statesinD pref withd pref(s)>0, the empirical preference-pair distribution in statesfollows
ρs
pref(ah, an) = E(s′,a+,a−)∼Dpref I[s′ =s, a + =a h, a− =a n]
E(s′,a+,a−)∼Dpref I[s′ =s] .(5)
The ideal preference-pair distribution at any state s in Dpref is simply the joint distribution of (ah, an):
ρs
ideal(ah, an) =π h(ah |s)π n(an |s)onA × A.
Finally, we define the optimization error of the agent policyπn as ϵ=L pref(πn)−Lpref(πh). We recall
that Lpref(π) =− E
(s,a+,a−)∼Dpref
[logσ(βlogπ(a + |s)−βlogπ(a − |s))] , where β is a positive
constant andσ(·)is the Sigmoid function. Under these notations, we have the following Thm. 4.1.
Theorem 4.1.We denote the Q-function of the human policy πh as Q∗(s, a). We assume that for
any (s, a, a′), |Q∗(s, a)−Q ∗(s, a′)| ≤U , |logπ h(a|s)−logπ h(a′|s)| ≤M , and |logπ n(a|s)−
logπ n(a′|s)| ≤M, whereU, M >0are constants. Whenβis small enough, we have
J(π h)−J(π n) =O(
p
ϵ+δ pref +δ dist).(6)
Here we explain the insights of Thm. 4.1 as follows. In our choice of the preference horizon L,
the key is to balance the two error terms δdist and δpref. Recall that the distribution shift term δdist
measures how close the state distributions are when there is no human intervention ( dπn) and the
state distribution represented in the preference dataset (dpref). Increasing L decreases δdist because
the preference dataset will contain more predicted states ˜si from the agent’s future trajectories. In
contrast, the preference error term δpref captures the misalignment between the true but unobserved
human action at a future state t′ > t and the bootstrapped corrective action ah at step t, which we
assume would also apply at t′. Therefore, the longer the preference horizon, the larger δpref, because
the difference between the human actions ah in state s and the predicted ˜sL grows as L increases. In
Fig. 8, we visualize the effects ofLon the performance of PPL. We prove Thm. 4.1 in Appendix F.
4.3 Implementation Details
Tasks.As shown in Fig. 4, we conduct experiments on control tasks and manipulation tasks with
different observation and action spaces. For the control task, we consider the MetaDrive driving
6

## Page 7

experiments [16], where the agent must navigate towards the destination in heavy-traffic scenes
without crashing into obstacles or other vehicles. The agent uses the sensory state vector s∈R 259
as its observation and outputs a control signal a= (a 0, a1)∈[−1,1] 2 representing the steering
angle and the acceleration, respectively. We evaluate the agent’s learned policy in a held-out test
environment separate from the training environments.
For manipulation tasks, we consider the Table Wiping and Nut Assembly tasks from the Robosuite
environment [49]. In the Table Wiping task, the robot arm must learn to wipe the whiteboard surface
and clean all of the markings. The positions of these markings are randomized at the beginning of
each episode. The states are s∈R 34 and actions are a∈R 6 (3 translations in the XYZ axes and 3
rotations around the XYZ axes). In the Nut Assembly task, the robot must grab a metal ring from
a random initial pose and place it over a target cylinder at a fixed location. The states are s∈R 51
and actions are a∈R 7, where the additional dimension in the action space represents opening or
closing the gripper. In both manipulation tasks, the simulated UR5e robot arm uses fixed-impedance
operational-space control to achieve the commanded pose.
Trajectory Prediction Model.In PPL, we need to predict the future states f(s, a n, H) =
(s,˜st+1,· · ·,˜s t+H ) from the current state s. We implement f by running an H-step simulator rollout
from the current state s, repeatedly applying action an to collect the sequence (˜st+1,· · ·,˜s t+H ).
ThisH-step simulator rollout runs at up to 1,000 fps on a CPU.
In real-world tasks such as autonomous driving, simulator rollouts often deviate from reality because
vehicle dynamics parameters are imperfect and other traffic participants behave unpredictably. To
predict future motion with minimal overhead, prior work directly propagates the ego-vehicle’s
state through a physics model [18, 29, 13]. Following this approach, we use the kinematic bicycle
model [30] to simulate H= 10 steps, assuming all other traffic participants remain stationary.
Compared with the data-driven approaches [50, 5, 21], this rule-based predictor requires only forward
integration of a single vehicle and produces short-term trajectories whose accuracy closely matches
simulator rollouts. This lightweight extrapolation method runs at about 3,000 fps on a CPU, enabling
real-time prediction with minimal overhead. Our ablation studies confirm that replacing the simulator
with our bicycle-model predictions incurs negligible performance loss (Table 2, rows 9-10).
5 Experiments
5.1 Experimental Setting
Neural Policies as Proxy Human Policies.Experiments with real human participants are time-
consuming and exhibit high variability between trials. Following the prior works on interactive
imitation learning [10, 27], in addition to real-human experiments, we also incorporate neural policies
in the training loop of PPL to approximate human policies in Table 3, 4, and 5. The neural experts
are trained using PPO-Lagrangian [33] for 20 million environment steps.
In MetaDrive, the neural expert uses the following takeover rule when training all baselines and
our method PPL: if the predicted trajectory τ=f(s, a n, H) contains any safety violation, such as
crashes or going off the road, or the average speed is too slow, the expert takes control for the next
H steps. In RoboSuite, the neural expert intervenes when the cumulative reward over the predicted
trajectory τ falls below a threshold ϵ. We set ϵ= 1 for the Table Wiping task and ϵ= 2 for the Nut
Assembly task.
In Table 1, we report experiments involving real humans in the MetaDrive safety benchmark. In
Table 3, Table 4, and Table 5, we report experiments with the neural policy as the proxy human policy
in the MetaDrive, Table Wiping, and Nut Assembly tasks, respectively.
Evaluation Metrics.In the Table Wiping task and Nut Assembly task, we report thesuccess rate,
the ratio of episodes where the agent reaches the destination. In the MetaDrive safety benchmark, we
also report theepisodic returnandroute completion rateduring evaluation. The route completion
rate is the ratio of the agent’s successfully traveled distance to the length of the complete route.
We train each interactive imitation learning baseline five times using distinct random seeds. Then, we
roll out 50 trajectories generated by each model in the held-out evaluation environment and average
each evaluation metric as the model’s performance. During the evaluation, no expert is involved. The
standard deviation is provided. We fix H= 10 for all the interactive imitation learning baselines. In
7

## Page 8

(a) MetaDrive(b) Robosuite: Table Wiping(c) Robosuite: Nut Assembly
Figure 4: Human interfaces of the three tasks: MetaDrive (a), Table Wiping (b), and Nut Assembly
(c). In (a), the agent’s forecasted trajectory (the red dots) leads to a collision, prompting the expert to
intervene via the gamepad (blue dots show the predicted rollout of the expert). In (b) and (c), the
expert observes the agent’s forecasted trajectory and intervenes via the keyboard if necessary.
Figure 5: Training process of PPL in the MetaDrive environment with the human expert over 20K
steps. We plot the test success rate (left), training takeover rate (top right), and training episodic
safety cost (bottom right). During training, when the agent’s forecasted trajectory (red dots) leads to
a collision, the human expert intervenes via the gamepad, and the corrected rollout is shown (blue
dots). When the agent’s forecasted trajectory is safe, it is visualized in green dots. The agent becomes
autonomous and performant during training, requiring fewer human interventions to maintain safety.
PPL, we fix β= 0.1 , choose L= 4 for the MetaDrive benchmark and Table Wiping task, and set
L= 6 for the Nut Assembly task. In Fig. 8, we show how the choice of L affects the performance of
PPL in the MetaDrive benchmark.
We also report the total number of human-involved transitions (human data usage) and theoverall
intervention rate, which is the ratio of human data usage to total data usage. These show how much
effort humans make to teach the agents.
Human Interfaces.Human subjects can take control through the Xbox Wireless Controller or the
keyboard and monitor the training process by visualizing environments on the screen. The predicted
trajectories are updated every H= 10 steps (one second), so that the human expert can intervene
promptly before the agent causes any safety violations and undesired behaviors.
Baselines.We test two imitation learning baselines: Behavior Cloning (BC) and GAIL [ 11], and
two confidence-based IIL methods: Ensemble-DAgger [23] and Thrifty-DAgger [12]. Four human-
in-the-loop IIL methods that learn from active human involvement are tested: Intervention Weighted
Regression (IWR) [22], Human-AI Copilot Optimization (HACO) [17], Expert Intervention Learning
(EIL) [42], and Proxy Value Propagation [28].
5.2 Baseline Comparison
In Table 1, we report the performance of our PPL and all the baselines with real human experts in the
MetaDrive safety benchmark. Our method PPL outperforms all the baselines and achieves a success
rate of 76% within 10K steps. The whole experiment of PPL takes only 12 minutes on a desktop
computer with an Nvidia GeForce RTX 4080 GPU.
In Table 3, 4, and 5, we report the performance of our PPL and all the baselines with neural experts
as proxy human policies in MetaDrive, Table Wiping, and Nut Assembly tasks, respectively. We
also plot the curves of the test-time success rate in Fig. 3. These tables and Fig. 3 show that PPL
achieves both fewer expert data usage and environment samples needed in both driving tasks and
8

## Page 9

Table 1: Comparison of methods with training/testing statistics in the MetaDrive environment with
the real human expert. The overall intervention rate is given together with the human data usage.
Method Human-in-the-Loop Training Testing
Human Data Usage Total Data Usage Success Rate Episodic Return Route Completion
Human Expert – 20K – 0.95±0.04349.2±18.20.98±0.01
BC✗20K – 0.0±0.053.5±22.80.16±0.07
GAIL✗20K 1M 0.14±0.03146.2±17.10.44±0.05
Ensemble-DAgger✓3.8K (0.38) 10K 0.36±0.11233.8±21.30.70±0.02
Thrifty-DAgger✓3.2K (0.32) 10K 0.45±0.04221.5±26.40.62±0.04
PVP✓4.9K (0.49) 10K 0.46±0.08267.3±15.00.71±0.04
IWR✓5.2K (0.52) 10K 0.23±0.10246.7±10.70.62±0.02
EIL✓6.9K (0.69) 10K 0.01±0.01137.3±26.10.40±0.08
HACO✓6.3K (0.63) 10K 0.11±0.05154.7±14.70.45±0.09
PPL (Ours)✓2.9K(0.29) 10K0.76±0.07324.8±9.20.90±0.06
Table 2: Ablation studies in MetaDrive with 10K
total data usage. We use the neural expert as the
proxy human policy.
Method
Expert Data
Usage
Route
Completion
Success
Rate
Imitation ona + 1.9K 0.65 0.36
PPL with randoma + 2.2K 0.73 0.45
PPL with randoma − 2.3K 0.69 0.38
PPL with DPO1.6K 0.91 0.80
PPL with IPO 2.6K 0.61 0.35
PPL with SLiC-HF 3.0K 0.59 0.32
PPL with BC loss only 2.0K 0.72 0.42
PPL with CPO loss only 5.8K 0.31 0.04
PPL with rule-basedf1.9K 0.91 0.78
PPL (Ours)1.8K 0.92 0.81
PVP
PPL (Ours)
Figure 6: We plot the steering control sequences
for both PVP and PPL on the same MetaDrive
map, with arrows representing the steering an-
gles every five steps. Both agents are trained to
10K steps. Compared to PVP, our method yields
smoother steering and more consistent speeds,
especially when navigating close to obstacles.
robot manipulation tasks while significantly outperforming baselines in testing performance. These
results suggest that our construction of the preference dataset accurately reflects human preferences
and helps speed up imitation learning. In addition, Fig. 6 shows that our method PPL produces
smoother control sequences and generates trajectories that better align with human preferences.
5.3 Ablation Studies
In Table 2, we perform ablation studies of our PPL in the MetaDrive safety benchmark with the
neural expert as proxy human policies.
Discarding positive or negative actions:In the first three rows of Table 2, we show that the advantage
of our method PPL arises from the constructed preference pairs (˜s, a+, a−) in the preference data
Dpref (Fig. 2 (right)), instead of merely emulating the positive actions a+ or simply avoiding taking
the negative actions a− in the preference buffer. As shown in Table 2, discarding the negative
actions a− and performing Behavior Cloning on the positive actions (Imitation on a+) leads to poor
performance, which is even worse than directly imitating the expert demonstrations in the human
buffer Dh (PPL with BC loss only). In addition, replacing the positive actions by random actions
(PPL with random a+) or the negative actions by random actions (PPL with random a−) also fails to
solve the MetaDrive benchmark.
Preference-based RL objectives:In our learning objective Eq. 4, we use the Contrastive Preference
Optimization (CPO) loss [45] to learn from the preference dataset Dpref. In Table 2 (rows 4–6), we
also report the performance of using other preference-based RL objectives from Direct Preference
Optimization (DPO) [31], IPO [1], and SLiC-HF [48]. For DPO and IPO, we use a reference policy
trained by Behavior Cloning from 10K expert demonstrations. Table 2 shows that using IPO (PPL
with IPO) and SLiC-HF (PPL with SLiC-HF) objectives degrade the performance of PPL. Using the
DPO objective (PPL with DPO) does not hurt the performance of PPL. However, the DPO objective
requires access to a pretrained reference policy, while our learning objective Eq. 4 does not.
Discarding the BC loss or preference loss:As shown in row 7 of Table 2, discarding the CPO loss
Lpref in Eq. 4 (PPL with BC loss only) significantly damages the performance of PPL. Discarding
9

## Page 10

0.00 0.05 0.10 0.15 0.20 0.25 0.30
Trajectory Prediction Error 
0.50
0.55
0.60
0.65
0.70
0.75
0.80T est Success Rate
PPL's Performance vs. Prediction Noise
PPL (Ours)
PVP
Figure 7: Performance of PPL under varying
trajectory-prediction noise levels ϵ in MetaDrive.
PPL still outperforms PVP when the trajectory
predictor is imperfect.
2 3 4 5 6
Preference Horizon L
0.55
0.60
0.65
0.70
0.75
0.80T est Success Rate
PPL's Performance vs. Preference Horizon
PPL (10K steps)
PPL (15K steps)
PVP (15K steps)
Figure 8: Performance of PPL with different
preference horizons L in MetaDrive with 10K
and 15K total data usage. PPL has the best learn-
ing efficiency when we set L= 4 in MetaDrive.
the BC loss (PPL with CPO loss only) also damages the performance, because the BC loss helps
regularize our learned policy and avoid it deviating too much from the expert demonstrations.
Rule-based trajectory prediction model:Following Sec. 4.3, we also implement a rule-based
trajectory prediction model f by simulating the ego-vehicle dynamics for H steps. Using a rule-based
f (PPL with rule-based f) has negligible effects on the performance of PPL. This shows that our
method still outperforms the IIL baselines even without relying on simulator rollouts.
5.4 Robustness Analysis
In Sec. 5.4, we evaluate PPL’s robustness to noise in the trajectory predictor (Fig. 7). We also
visualize the effect of the preference horizonLon PPL in Fig. 8.
In Fig. 7, we show that PPL is robust to noise in trajectory predictors. With an imperfect predictive
model, PPL still outperforms all the baselines. We inject random Gaussian noiseenoise to the outputs ˜s
of the trajectory predictor, and we set the norm ||enoise||2 =ϵ∗ ||˜s||2. Then we gradually increase the
constant ϵ to test PPL’s robustness to noises in trajectory predictors. We use MetaDrive, Table Wiping,
and Nut Assembly environments following the same setups from Tables 3, 4, and 5, respectively. We
find that with a noisy predictive model, PPL still outperforms all the baselines in MetaDrive and Table
Wiping when the noise ϵ≤0.25 . In Nut Assembly, PPL outperforms the baselines when ϵ≤0.125 .
In Fig. 8, we visualize how the preference horizon L affects the test success rate of PPL in the
MetaDrive safety benchmark with 10K and 15K total data usage. As L increases from 2 to 4, the
agent gains additional corrective information from forecasted states in the preference buffer and
achieves higher success rates. Beyond L= 4 , however, the benefit tapers off and eventually degrades,
since overly long horizons yield less accurate preference labels. Therefore, we observe peak learning
efficiency at L= 4 . Notably, when 3≤L≤5 , PPL trained for only 10K steps already outperforms
PVP trained for 15K steps. With an appropriately chosen preference horizon, PPL can substantially
reduce both training time and expert monitoring effort.
6 Conclusion
In this work, we propose Predictive Preference Learning from Human Interventions (PPL), a novel
interactive imitation learning algorithm that applies preference learning over predicted future trajecto-
ries to capture implicit human preferences. By converting each expert intervention into contrastive
preference labels across forecasted states, PPL directs corrective feedback toward the regions of
the state space the agent is most likely to explore. This approach substantially improves learning
efficiency and reduces both the number of required demonstrations and the expert’s cognitive load,
without offline pretraining and reward engineering.
Limitations.We assume that the expert always knows the optimal corrective action and demon-
strates it accurately, whereas human demonstrations can be suboptimal or inconsistent. Additionally,
all our experiments are conducted in simulation. The effectiveness and safety of PPL on real robots
operating in physical environments remain to be explored in future works.
10

## Page 11

Acknowledgment: This work was supported by the NSF Grants CCF-2344955 and IIS-2339769,
and ONR grant N000142512166. The human experiment in this study is approved through the
IRB#23-000116 at UCLA. ZP was supported by the Amazon Fellowship via UCLA Science Hub.
References
[1] Mohammad Gheshlaghi Azar, Zhaohan Daniel Guo, Bilal Piot, Remi Munos, Mark Rowland, Michal
Valko, and Daniele Calandriello. A general theoretical paradigm to understand learning from human
preferences. InInternational Conference on Artificial Intelligence and Statistics, pages 4447–4455. PMLR,
2024.
[2] Haoyuan Cai, Zhenghao Peng, and Bolei Zhou. Robot-gated interactive imitation learning with adaptive
intervention mechanism.International Conference on Machine Learning, 2025.
[3] Sonia Chernova and Andrea L Thomaz.Robot learning from human teachers. Springer Nature, 2022.
[4] Sonia Chernova and Manuela Veloso. Interactive policy learning through confidence-based autonomy.
Journal of Artificial Intelligence Research, 34:1–25, 2009.
[5] Cheng Chi, Zhenjia Xu, Siyuan Feng, Eric Cousineau, Yilun Du, Benjamin Burchfiel, Russ Tedrake, and
Shuran Song. Diffusion policy: Visuomotor policy learning via action diffusion.The International Journal
of Robotics Research, page 02783649241273668, 2023.
[6] Paul F. Christiano, Jan Leike, Tom B. Brown, Miljan Martic, Shane Legg, and Dario Amodei. Deep
reinforcement learning from human preferences. In Isabelle Guyon, Ulrike von Luxburg, Samy Bengio,
Hanna M. Wallach, Rob Fergus, S. V . N. Vishwanathan, and Roman Garnett, editors,Advances in Neural
Information Processing Systems 30: Annual Conference on Neural Information Processing Systems 2017,
December 4-9, 2017, Long Beach, CA, USA, pages 4299–4307, 2017.
[7] Bin Fang, Shidong Jia, Di Guo, Muhua Xu, Shuhuan Wen, and Fuchun Sun. Survey of imitation learning
for robotic manipulation.International Journal of Intelligent Robotics and Applications, 3:362–369, 2019.
[8] Aditya Ganapathi, Priya Sundaresan, Brijen Thananjeyan, Ashwin Balakrishna, Daniel Seita, Jennifer
Grannen, Minho Hwang, Ryan Hoque, Joseph E Gonzalez, Nawid Jamali, et al. Learning dense visual
correspondences in simulation to smooth and fold real fabrics. In2021 IEEE International Conference on
Robotics and Automation (ICRA), pages 11515–11522. IEEE, 2021.
[9] Lin Guan, Mudit Verma, Sihang Guo, Ruohan Zhang, and Subbarao Kambhampati. Widening the pipeline
in human-guided reinforcement learning with explanation and context-aware data augmentation.Advances
in Neural Information Processing Systems, 34, 2021.
[10] Joey Hejna, Rafael Rafailov, Harshit Sikchi, Chelsea Finn, Scott Niekum, W Bradley Knox, and Dorsa
Sadigh. Contrastive preference learning: learning from human feedback without rl.arXiv preprint
arXiv:2310.13639, 2023.
[11] Jonathan Ho and Stefano Ermon. Generative adversarial imitation learning. In Daniel D. Lee, Masashi
Sugiyama, Ulrike von Luxburg, Isabelle Guyon, and Roman Garnett, editors,Advances in Neural In-
formation Processing Systems 29: Annual Conference on Neural Information Processing Systems 2016,
December 5-10, 2016, Barcelona, Spain, pages 4565–4573, 2016.
[12] Ryan Hoque, Ashwin Balakrishna, Ellen Novoseller, Albert Wilcox, Daniel S. Brown, and Ken Goldberg.
Thriftydagger: Budget-aware novelty and risk gating for interactive imitation learning, 2021.
[13] Nico Kaempchen, Bruno Schiele, and Klaus Dietmayer. Situation assessment of an autonomous emergency
brake for arbitrary vehicle-to-vehicle collision scenarios.IEEE Transactions on Intelligent Transportation
Systems, 10(4):678–687, 2009.
[14] Sham M. Kakade and John Langford. Approximately optimal approximate reinforcement learning. In
Claude Sammut and Achim G. Hoffmann, editors,Machine Learning, Proceedings of the Nineteenth
International Conference (ICML 2002), University of New South Wales, Sydney, Australia, July 8-12, 2002,
pages 267–274. Morgan Kaufmann, 2002.
[15] Michael Kelly, Chelsea Sidrane, Katherine Driggs-Campbell, and Mykel J Kochenderfer. Hg-dagger:
Interactive imitation learning with human experts. In2019 International Conference on Robotics and
Automation (ICRA), pages 8077–8083. IEEE, 2019.
11

## Page 12

[16] Quanyi Li, Zhenghao Peng, Lan Feng, Qihang Zhang, Zhenghai Xue, and Bolei Zhou. Metadrive:
Composing diverse driving scenarios for generalizable reinforcement learning.IEEE transactions on
pattern analysis and machine intelligence, 2022.
[17] Quanyi Li, Zhenghao Peng, and Bolei Zhou. Efficient learning of safe driving policy via human-ai copilot
optimization. InInternational Conference on Learning Representations, 2022.
[18] Chiu-Feng Lin, A Galip Ulsoy, and David J LeBlanc. Vehicle dynamics and external disturbance estimation
for vehicle path prediction.IEEE Transactions on Control Systems Technology, 8(3):508–518, 2000.
[19] Huihan Liu, Soroush Nasiriany, Lance Zhang, Zhiyao Bao, and Yuke Zhu. Robot learning on the job:
Human-in-the-loop autonomy and learning during deployment.The International Journal of Robotics
Research, page 02783649241273901, 2022.
[20] Huihan Liu, Yu Zhang, Vaarij Betala, Evan Zhang, James Liu, Crystal Ding, and Yuke Zhu. Multi-task
interactive robot fleet learning with visual world models.arXiv preprint arXiv:2410.22689, 2024.
[21] Yicheng Liu, Jinghuai Zhang, Liangji Fang, Qinhong Jiang, and Bolei Zhou. Multimodal motion prediction
with stacked transformers. InProceedings of the IEEE/CVF conference on computer vision and pattern
recognition, pages 7577–7586, 2021.
[22] Ajay Mandlekar, Danfei Xu, Roberto Martín-Martín, Yuke Zhu, Li Fei-Fei, and Silvio Savarese. Human-
in-the-loop imitation learning using remote teleoperation.ArXiv preprint, abs/2012.06733, 2020.
[23] Kunal Menda, Katherine Driggs-Campbell, and Mykel J Kochenderfer. Ensembledagger: A bayesian
approach to safe imitation learning. In2019 IEEE/RSJ International Conference on Intelligent Robots and
Systems (IROS), pages 5041–5048. IEEE, 2019.
[24] Yu Meng, Mengzhou Xia, and Danqi Chen. Simpo: Simple preference optimization with a reference-free
reward.Advances in Neural Information Processing Systems, 37:124198–124235, 2024.
[25] Long Ouyang, Jeff Wu, Xu Jiang, Diogo Almeida, Carroll L Wainwright, Pamela Mishkin, Chong Zhang,
Sandhini Agarwal, Katarina Slama, Alex Ray, et al. Training language models to follow instructions with
human feedback.arXiv preprint arXiv:2203.02155, 2022.
[26] Malayandi Palan, Gleb Shevchuk, Nicholas Charles Landolfi, and Dorsa Sadigh. Learning reward functions
by integrating human demonstrations and preferences. InRobotics: Science and Systems, 2019.
[27] Zhenghao Peng, Quanyi Li, Chunxiao Liu, and Bolei Zhou. Safe driving via expert guided policy
optimization. In5th Annual Conference on Robot Learning, 2021.
[28] Zhenghao Mark Peng, Wenjie Mo, Chenda Duan, Quanyi Li, and Bolei Zhou. Learning from active human
involvement through proxy value propagation.Advances in neural information processing systems, 36,
2024.
[29] Romain Pepy, Alain Lambert, and Hugues Mounier. Reducing navigation errors by planning with realistic
vehicle model. In2006 IEEE Intelligent Vehicles Symposium, pages 300–307. IEEE, 2006.
[30] Philip Polack, Florent Altché, Brigitte d’Andréa Novel, and Arnaud de La Fortelle. The kinematic bicycle
model: A consistent model for planning feasible trajectories for autonomous vehicles? In2017 IEEE
intelligent vehicles symposium (IV), pages 812–818. IEEE, 2017.
[31] Rafael Rafailov, Archit Sharma, Eric Mitchell, Christopher D Manning, Stefano Ermon, and Chelsea Finn.
Direct preference optimization: Your language model is secretly a reward model.Advances in Neural
Information Processing Systems, 36:53728–53741, 2023.
[32] Harish Ravichandar, Athanasios S Polydoros, Sonia Chernova, and Aude Billard. Recent advances in robot
learning from demonstration.Annual Review of Control, Robotics, and Autonomous Systems, 3:297–330,
2020.
[33] Alex Ray, Joshua Achiam, and Dario Amodei. Benchmarking safe exploration in deep reinforcement
learning.arXiv preprint arXiv:1910.01708, 7(1):2, 2019.
[34] Siddharth Reddy, Anca D Dragan, and Sergey Levine. Shared autonomy via deep reinforcement learning.
Robotics: Science and Systems, 2018.
[35] Stéphane Ross and Drew Bagnell. Efficient reductions for imitation learning. InProceedings of the
thirteenth international conference on artificial intelligence and statistics, pages 661–668. JMLR Workshop
and Conference Proceedings, 2010.
12

## Page 13

[36] Stéphane Ross, Geoffrey Gordon, and Drew Bagnell. A reduction of imitation learning and structured
prediction to no-regret online learning. InProceedings of the fourteenth international conference on
artificial intelligence and statistics, pages 627–635. JMLR Workshop and Conference Proceedings, 2011.
[37] Dorsa Sadigh, Anca D Dragan, Shankar Sastry, and Sanjit A Seshia. Active preference-based learning of
reward functions.UC Berkeley, 2017.
[38] H. Saeidi, Justin D. Opfermann, Michael Kam, Sudarshan Raghunathan, S. Leonard, and A. Krieger. A
confidence-based shared control strategy for the smart tissue autonomous robot (star). In2018 IEEE/RSJ
International Conference on Intelligent Robots and Systems (IROS), pages 1268–1275, 2018.
[39] Mikayel Samvelyan, Tabish Rashid, Christian Schroeder De Witt, Gregory Farquhar, Nantas Nardelli,
Tim GJ Rudner, Chia-Man Hung, Philip HS Torr, Jakob Foerster, and Shimon Whiteson. The starcraft
multi-agent challenge.ArXiv preprint, abs/1902.04043, 2019.
[40] William Saunders, Girish Sastry, Andreas Stuhlmueller, and Owain Evans. Trial without error: Towards
safe reinforcement learning via human intervention. InProceedings of the 17th International Conference on
Autonomous Agents and MultiAgent Systems, pages 2067–2069. International Foundation for Autonomous
Agents and Multiagent Systems, 2018.
[41] Esmaeil Seraj, Kin Man Lee, Zulfiqar Zaidi, Qingyu Xiao, Zhaoxin Li, Arthur Nascimento, Sanne van
Waveren, Pradyumna Tambwekar, Rohan Paleja, Devleena Das, et al. Interactive and explainable robot
learning: A comprehensive review.Foundations and Trends® in Robotics, 12(2-3):75–349, 2024.
[42] Jonathan Spencer, Sanjiban Choudhury, Matthew Barnes, Matthew Schmittle, Mung Chiang, Peter Ra-
madge, and Siddhartha Srinivasa. Learning from interventions. InRobotics: Science and Systems (RSS),
2020.
[43] Nisan Stiennon, Long Ouyang, Jeffrey Wu, Daniel Ziegler, Ryan Lowe, Chelsea V oss, Alec Radford,
Dario Amodei, and Paul F Christiano. Learning to summarize with human feedback.Advances in neural
information processing systems, 33:3008–3021, 2020.
[44] Garrett Warnell, Nicholas R. Waytowich, Vernon Lawhern, and Peter Stone. Deep TAMER: interactive
agent shaping in high-dimensional state spaces. In Sheila A. McIlraith and Kilian Q. Weinberger, editors,
Proceedings of the Thirty-Second AAAI Conference on Artificial Intelligence, (AAAI-18), the 30th innovative
Applications of Artificial Intelligence (IAAI-18), and the 8th AAAI Symposium on Educational Advances in
Artificial Intelligence (EAAI-18), New Orleans, Louisiana, USA, February 2-7, 2018, pages 1545–1554.
AAAI Press, 2018.
[45] Haoran Xu, Amr Sharaf, Yunmo Chen, Weiting Tan, Lingfeng Shen, Benjamin Van Durme, Kenton Murray,
and Young Jin Kim. Contrastive preference optimization: Pushing the boundaries of llm performance in
machine translation.arXiv preprint arXiv:2401.08417, 2024.
[46] Zhenghai Xue, Zhenghao Peng, Quanyi Li, Zhihan Liu, and Bolei Zhou. Guarded policy optimization with
imperfect online demonstrations.arXiv preprint arXiv:2303.01728, 2023.
[47] Maryam Zare, Parham M Kebria, Abbas Khosravi, and Saeid Nahavandi. A survey of imitation learning:
Algorithms, recent developments, and challenges.IEEE Transactions on Cybernetics, 2024.
[48] Yao Zhao, Rishabh Joshi, Tianqi Liu, Misha Khalman, Mohammad Saleh, and Peter J Liu. Slic-hf:
Sequence likelihood calibration with human feedback.arXiv preprint arXiv:2305.10425, 2023.
[49] Yuke Zhu, Josiah Wong, Ajay Mandlekar, Roberto Martín-Martín, Abhishek Joshi, Soroush Nasiriany, and
Yifeng Zhu. robosuite: A modular simulation framework and benchmark for robot learning.arXiv preprint
arXiv:2009.12293, 2020.
[50] Alex Zyner, Stewart Worrall, and Eduardo Nebot. A recurrent neural network solution for predicting driver
intention at unsignalized intersections.IEEE Robotics and Automation Letters, 3(3):1759–1764, 2018.
13

## Page 14

A Algorithm
We summarize our method PPL in Alg. 1.
Algorithm 1Predictive Preference Learning from Human Interventions (PPL)
1:Input:HyperparametersH, L, β.
2:fortimestepk= 0, H,2H, . . .do
3:Agent samples actiona n ∼π n(sk).
4:Predict future trajectoryτ=f(s k, an, H) = (s k,˜sk+1,· · ·,˜s k+H ).
5:Human observesτto decide whether to take over in the nextHsteps.
6:fortimestept=k, k+ 1,· · ·, k+H−1do
7:ifHuman takes overthen
8:Human takes actiona h ∼π h(st).
9:Add(s t, ah)to the human bufferD h.
10:Agent samples actiona n ∼π n(st).
11:Predict future trajectoryτ ′ =f(s t, an, L) = (s t,˜st+1,· · ·,˜s t+L).
12:Add(˜s, a h, an)to the preference datasetD pref for each˜sin(s t,˜st+1,· · ·,˜s t+L).
13:Observes t+1 ∼ P(· |s t, ah).
14:else
15:Agent samples actiona n ∼π n(st).
16:Observes t+1 ∼ P(· |s t, an).
17:end if
18:Train policyπ n with loss function Eq. 4.
19:end for
20:end for
21:Output:Policyπ n.
B Additional Experimental Results
We report the performance of our PPL and all the baselines with neural experts as proxy human poli-
cies in MetaDrive (Table 3), Table Wiping (Table 4), and Nut Assembly (Table 5) tasks, respectively.
The test success rate curves of all three tasks are shown in Fig. 3.
Table 3: Comparison of methods with training/testing statistics in the MetaDrive environment with
the neural expert as the proxy human policy. The overall intervention rate is given together with the
expert data usage.
Method Expert-in-the-Loop Training Testing
Expert Data Usage Total Data Usage Success Rate Episodic Return Route Completion
Neural Expert – – – 0.83±0.07340.2±15.90.93±0.02
BC✗20K – 0.12±0.04142.7±27.50.46±0.07
GAIL✗20K 1M 0.34±0.08196.5±14.10.60±0.09
Ensemble-DAgger✓3.2K (0.32) 10K 0.41±0.08238.6±13.00.69±0.07
Thrifty-DAgger✓2.9K (0.29) 10K 0.49±0.07248.2±27.80.75±0.06
PVP✓2.5K (0.25) 10K 0.56±0.07258.1±23.40.76±0.05
IWR✓2.7K (0.27) 10K 0.33±0.11217.0±20.90.67±0.06
EIL✓3.9K (0.39) 10K 0.11±0.06131.8±29.50.42±0.11
HACO✓2.6K (0.26) 10K 0.36±0.15210.2±25.20.64±0.10
PPL (Ours)✓1.2K(0.20)6K 0.80±0.04329.9±13.40.92±0.03
Table 4: Results of different approaches in Table
Wiping.
Method Expert Data Usage Total Data Success Rate
Neural Expert – – 0.84
BC 10K – 0.11
GAIL 10K 1M 0.37
PVP 2.3K 4K 0.58
IWR 2.5K 4K 0.51
EIL 2.4K 4K 0.53
HACO 2.9K 4K 0.48
PPL (Ours)0.2K 2K 0.80
Table 5: Results of different approaches in Nut
Assembly.
Method Expert Data Usage Total Data Success Rate
Neural Expert – – 0.60
BC 100K – 0.02
GAIL 100K 1M 0.08
PVP 49K 200K 0.35
IWR 54K 200K 0.29
EIL 48K 200K 0.25
HACO 77K 200K 0.15
PPL (Ours)48K 200K 0.51
14

## Page 15

The neural experts in Table 3, 4, and 5 are trained with PPO-Lagrangian [33] for 20M environment
steps, yet their test success rates still fall short of 100% for the following reasons. The MetaDrive
safety environments occasionally generate rare but challenging scenarios that even a well-trained
policy may fail to handle. In the Table Wiping task, the neural expert sometimes fails to remove one
or two markings on the whiteboard, leaving a small patch of dirt uncleaned. In the Nut Assembly
task, successful grasping requires the gripper to be precisely aligned with the metal ring’s handle,
which is highly sensitive to even minor action errors.
C Demo Video
We have attached our demo video to https://metadriverse.github.io/ppl. This video shows
how we conduct human experiments and the evaluation results of our method Predictive Preference
Learning from Human Interventions (PPL). This video includes five sections:
1. The first section gives an overview of Predictive Preference Learning, showing what human
observes on the screen and how human provides corrective demonstrations in an episode.
2. The second section is the footage of the MetaDrive human experiment, where the human
expert interacts with the driving agent via a gamepad.
3. The third section shows the evaluation results of the PPL agent in a held-out MetaDrive
test environment. We compare our approach PPL with the PVP baseline [ 28], and both
agents are trained to 10K timesteps. The evaluation results show that our approach PPL has
a higher test success rate and lower safety cost.
4. The fourth section shows the applicability of our methods to manipulation tasks in Ro-
bosuite [49]: Table Wiping and Nut Assembly. PPL successfully imitates the expert and
accomplishes both tasks in evaluation environments.
5. In the fifth section, we provide a full training session on MetaDrive. The video is played at
5× speed, and it shows how a human expert trains a PPL agent on MetaDrive in under 12
minutes, with approximately 1.8K demonstration steps and 10K environment steps.
D Human Subject Research Protocol
Human Participants.Five university students (ages 20–30) with valid driver’s licenses and video
gaming experience took part in the study voluntarily. After receiving a detailed overview of the
procedures and providing written informed consent under an IRB-approved protocol, each participant
completed a hands-on familiarization session. During this session, they were informed how the
predicted trajectories were shown on screen, and they practiced with our control interface and learning
environments until performing ten consecutive successful runs before the main experiments.
Main Experiment.Each participant began with one or two fully manual episodes to build con-
fidence, and then ceded control to the agent when they felt safe. Participants were instructed to
intervene only when the agent’s predicted trajectory appeared unsafe, illegal, or inconsistent with
their desired actions. They were directed to prioritize safe task completion and then to guide the
agent toward their personal driving or manipulation preferences.
E Notations
Before we prove Theorem 4.1, we recall all the notations in this work. We denote the human
policy as πh and the novice policy as πn. For any stochastic policy π(a|s) and the initial state
distribution d0 on state space S, we define the value function J(π) as the expected cumulative
return: J(π) = E
τ∼P π
[
∞P
t=0
γtr(st, at)], wherein τ= (s 0, a0, s1, a1, ...) is the trajectory sampled from
trajectory distribution Pπ induced by π, d0 and the state transition distribution P. We also denote
the Q-function of policy π as Q(s, a) = E
τ∼P π
[
∞P
t=0
γtr(st, at)
s0 =s, a 0 =a]. And we define the
discounted state distribution underπasd π(s) = (1−γ) E
τ∼P π
[
∞P
t=0
γtI[st =s]].
15

## Page 16

In our algorithm PPL, we have a preference dataset Dpref containing preference pairs (s, a+, a−).
The preference loss function of policyπin PPL is defined as
Lpref(π) =|D pref|−1 X
(s′,a+,a−)∈Dpref

−logσ
 
βlogπ(a + |s)−βlogπ(a − |s)

,(7)
whereβis a positive constant andσ(x) = (1 + exp(−x)) −1 is the Sigmoid function.
We denote the state distribution inD pref as
dpref(s) =|D pref|−1 X
(s′,a+,a−)∈Dpref
I[s′ =s].(8)
In addition, for any state s in Dpref with dpref(s)>0 , we denote the empirical preference-pair
distribution in statesas
ρs
pref(ah, an) =
P
(s′,a+,a−)∈Dpref
I[s′ =s, a + =a h, a− =a n]
P
(s′,a+,a−)∈Dpref
I[s′ =s] ,(9)
which is a distribution onA × A.
F Proof of Theorem 4.1
Our goal is to prove that the performance gap J(π h)−J(π n) between the human policy πh and the
agent policy πn can be bounded by the following three error terms: the state distribution shift δdist,
the quality of preference labels δpref, and the optimization error ϵ. We denote the total variation for
any two distributionsP, Qon the same space asD TV(P, Q) = 1
2 ∥P−Q∥ 1.
Here, we formally define the three error terms. The first state distribution shift error arises from
the difference between the distribution of states in the preference dataset Dpref (denoted as dpref(s))
and the discounted state distribution of the agent’s policy πn (denoted as dπn (s)). To define the
distribution shift errorδ dist in PPL, we use the total variation between the two distributions, i.e.,
δdist =D TV(dπn , dpref).(10)
The second error term arises from the misalignment of the positive actions in the preference dataset,
as the human action ah in each tuple (˜si, ah, an)∈ D pref is sampled in state s instead of the predicted
future state ˜si. In an ideal preference dataset, one would observe expert and novice actions drawn
directly at ˜si. To quantify this error, we define the following distribution ρs
ideal(ah, an) =π h(ah |
s)πn(an |s) on A × A for any state s, i.e., the distribution over pairs (ah, an) if both policies were
sampled at directly at states. Then we use
δpref = E
s∼dpref
DTV(ρs
ideal, ρs
pref)(11)
to define the errors in the preference dataset.
Finally, we define the optimization error of the agent policyπ n as
ϵ=L pref(πn)− L pref(πh).(12)
Under these notations, we have the following Thm. F.1. We note that when we choose a small
β≤M −2 (Mis defined in Thm. F.1), we have
J(π h)−J(π n) = 1
1−γ ·O
s
ϵ+ 4 log 2·δ pref
2β + 2δdist

. (13)
Theorem F .1(Formal Statement of Theorem 4.1).We denote the Q-function of the human policy
πh as Q∗(s, a). We assume that for any (s, a, a′), |Q∗(s, a)−Q ∗(s, a′)| ≤U , |logπ h(a|s)−
logπ h(a′|s)| ≤M , and |logπ n(a|s)−logπ n(a′|s)| ≤M , where U, M are positive constants.
Then, we have
J(π h)−J(π n)≤ U
1−γ ·
s
ϵ+ 4(βM+ log 2)·δ pref
2β + βM 2
8 + 2δdist

. (14)
16

## Page 17

Proof.The key is to combine Lem. F.2, Lem. F.3, and Lem. F.4 to obtain the bound.
From Lem. F.2, we can use the state distribution shift and the total variation of the two policies
πh, πn ond pref to bound the optimality gap:
J(π h)−J(π n)≤ U
1−γ ·

E
s∼dpref
DTV
 
πh(·|s), πn(·|s)

+ 2δdist

. (15)
In addition, for any policyπ, we define the function
g(π) = E
s∼dpref
E
a+∼πh(s),a−∼πn(s)

−logσ
 
βlogπ(a + |s)−βlogπ(a − |s)

,(16)
which represents the preference loss on ideal preference pairs, where a+, a− are sampled directly at
each states.
Using F.4, we can bound the total variation term E
s∼dpref
DTV
 
πh(·|s)

byg(π n)−g(π h):
E
s∼dpref
DTV
 
πh(·|s), πn(·|s)

≤
s
g(πn)−g(π h)
2β + βM 2
8 .(17)
In addition, by Lem. F.3, we can also bound g(πn)−g(π h) by the optimization error ϵ on Lpref and
the misalignment of preference levelsδ pref:
g(πn)−g(π h)≤ϵ+ 4(βM+ log 2)·δ pref .(18)
Combining Eq. 15, 17, and 18 yields Eq. 14.
Lemma F .2(Performance Optimality Gap on the State Distribution Shift).We recall that dpref (s) =
|Dpref|−1
E
(s′,a+,a−)∼Dpref
I[s′ =s], and we defineU= max
s∈S,a 1,a2∈A
|Q∗(s, a1)−Q ∗(s, a2)|.
Then, for any two stochastic policiesπ h, πn, we have
J(π h)−J(π n)≤ U
1−γ ·

E
s∼dpref
DTV
 
πh(·|s), πn(·|s)

+ 2δdist

. (19)
whereδ dist =D TV(dπn , dpref ).
Proof Sketch. The key is to use the Performance Difference Lemma (Lem. G.2) on J(π h)−J(π n),
yielding Eq. 20. Then, we can apply Lem. G.1, which bounds the expectation on s∼d pref
and s∼d πn by the distribution shift term δdist. Finally, applying the assumption U=
max
s∈S,a 1,a2∈A
|Q∗(s, a1)−Q ∗(s, a2)| bounds the difference of the Q-function by the total variation
betweenπ h andπ n.
Proof.By the Performance Difference Lemma (Lem. G.2), we have
J(π h)−J(π n) = 1
1−γ E
s∼dπn
E
ah∼πh(s),an∼πn(s)
[Q∗(s, ah)−Q ∗(s, an)]. (20)
By Lem. G.1, asd πn andd pref are two distributions on the same state spaceS, we have
(1−γ)(J(π h)−J(π n))
≤ E
s∼dpref
E
ah∼πh(s),an∼πn(s)
[Q∗(s, ah)−Q ∗(s, an)]
+ 2DTV(dπn , dpref )·max
s∈S
 E
ah∼πh(s),an∼πn(s)
[Q∗(s, ah)−Q ∗(s, an)]

≤ E
s∼dpref
E
ah∼πh(s),an∼πn(s)
[Q∗(s, ah)−Q ∗(s, an)]
+ 2δdist ·max
s∈S E
ah∼πh(s),an∼πn(s)
|Q∗(s, ah)−Q ∗(s, an)|
≤ E
s∼dpref

E
ah∼πh(s)
Q∗(s, ah)− E
an∼πn(s)
Q∗(s, an)

+ 2U·δ dist,
(21)
17

## Page 18

where we useU= max
s∈S,a 1,a2∈A
|Q∗(s, a1)−Q ∗(s, a2)|in the last inequality of Eq. 21.
In addition, πh(s) and πn(s) are two probability distributions on the same action space A. By Lem.
G.1, we have for anys∈ S,
E
ah∼πh(s)
Q∗(s, ah)− E
an∼πn(s)
Q∗(s, an)≤U·D TV
 
πh(·|s), πn(·|s)

.(22)
This proves that
J(π h)−J(π n)≤ U
1−γ ·

E
s∼dpref
DTV
 
πh(·|s), πn(·|s)

+ 2δdist

. (23)
Lemma F .3(Misalignment of Preference Pairs).We recall that the loss function of the policy π is
Lpref(π) =− E
(s,a+,a−)∼Dpref
[logσ(βlogπ(a + |s)−βlogπ(a − |s))] . And the optimization loss is
defined asϵ=L pref (πn)− L pref (πh).
In addition, following Eq. 16, for any policyπ, we define
g(π) = E
s∼dpref
E
a+∼πh(s),a−∼πn(s)

−logσ
 
βlogπ(a + |s)−βlogπ(a − |s)

.(24)
Under the assumption that for any (s, a, a′), |logπ h(a|s)−logπ h(a′|s)| ≤M , and |logπ n(a|s)−
logπ n(a′|s)| ≤M, we have
we can bound
g(πn)−g(π h)≤ϵ+ 4(βM+ log 2)·δ pref ,(25)
whereδ pref = E
s∼dpref
DTV(ρs
ideal, ρs
pref ).
Proof Sketch. The key is to apply Lem. G.1 on the two distributions ρs
ideal and ρs
pref, so that we can
bound the difference ofL pref(π)andg(π)for any policyπbyO(δ pref ).
Proof. For any s∈ S , we denote ρs
ideal(ah, an) =π h(ah |s)π n(an |s) , a probability distribution
on A×A . We also denote ρs
pref (ah, an) =ρ pref (s, ah, an)/dpref (s) for any s such that dpref (s)>0 ,
where we recall thatρ pref (s, ah, an) =|D pref|−1
E
(s′,a+,a−)∼Dpref
I[s′ =s, a + =a h, a− =a n].
The key is that ρs
ideal and ρs
pref are two distributions on the same space A × A, and we can apply
Lem. G.1 on Eq. 24 to obtain the proof.
We denotel π(s, a+, a−) =−logσ(βlogπ(a + |s)−βlogπ(a − |s)).
We also denotel π
max = max
s,a+,a−
|lπ(s, a+, a−)|, andl max = max(l πh
max, lπn
max). Then, for any policyπ,
g(π) = E
s∼dpref
E
a+∼πh(s),a−∼πn(s)
lπ(s, a+, a−)
= E
s∼dpref
E
(a+,a−)∼ρs
ideal
lπ(s, a+, a−)
≤ E
s∼dpref
h
2lπ
max ·D TV(ρs
ideal, ρs
pref ) + E
(a+,a−)∼ρs
pref
lπ(s, a+, a−)
i
= 2l π
maxδpref + E
s∼dpref
E
(a+,a−)∼ρs
pref
lπ(s, a+, a−)
= 2l π
maxδpref +|D pref|−1 X
(s,a+,a−)∈Dpref
lπ(s, a+, a−)
= 2l π
maxδpref +L pref(π).
(26)
Similarly, we can obtain thatg(π)≥ −2l π
maxδpref +L pref(π)for any policyπ. Thus we have
g(πn)−g(π h)≤4l maxδpref + (Lpref(πn)− L pref(πh)) = 4l maxδpref +ϵ.(27)
18

## Page 19

Finally, under the condition that |logπ(a|s)−logπ(a ′|s)| ≤M for any (s, a, a′), we have
|lπ(s, a, a′)| ≤ −logσ(−βM) = log(1 + exp(βM))≤βM+ log 2.
This implies thatl max ≤βM+ log 2and completes the proof.
Lemma F .4(Optimization Error Bounds the Total Variation).We assume that for any (s, a, a′),
|logπ h(a|s)−logπ h(a′|s)| ≤M, and|logπ n(a|s)−logπ n(a′|s)| ≤M.
We recall that g(π) = E
s∼dpref
E
a+∼πh(s),a−∼πn(s)
[−logσ(βlogπ(a + |s)−βlogπ(a − |s))] ,
which is defined in Eq. 16. Then we have
E
s∼dpref
DTV
 
πh(·|s), πn(·|s)

≤
s
g(πn)−g(π h)
2β + βM 2
8 .(28)
Proof Sketch.First, we define
f(π) =− β
2 E
s∼dpref
E
a+∼πh(s),a−∼πn(s)
h
logπ(a +|s)−logπ(a −|s)
i
+ log 2.(29)
Using the Taylor’s expansion on the function logσ(x) at x= 0 , when the policy π satisfies
|logπ(a|s)−logπ(a ′|s)| ≤Mfor any(s, a, a ′), we can obtain that|g(π)−f(π)| ≤ β2M 2
8 .
In addition, f(π n)−f(π h) bounds the KL divergence of the two policies πh and πn over s∼d pref.
So we can use Pinsker’s inequality to obtain the bound onDTV
 
πh(·|s), πn(·|s)

.
Proof. For any (s, a+, a−), we denote un(s, a+, a−) = logπ n(a+|s)−logπ n(a−|s), and
uh(s, a+, a−) = logπ h(a+|s)−logπ h(a−|s). From the assumptions, we can obtain that
|un(s, a+, a−)| ≤Mand|u h(s, a+, a−)| ≤M.
By definition of the function g(·), we have g(πn)−g(π h) =
E
s∼dpref
E
a+∼πh(s),a−∼πn(s)
[logσ(β·u h(s, a+, a−))−logσ(β·u n(s, a+, a−))].
The Taylor’s expansion oflogσ(x)atx= 0ensures that for anyx∈R, we have
logσ(x) + log 2− 1
2 x
 ≤ 1
8 x2.(30)
This ensures that
g(πn) = E
s∼dpref
E
a+∼πh(s),a−∼πn(s)

−logσ
 
β·u n(s, a+, a−)

≥log 2− β
2 E
s∼dpref
E
a+∼πh(s),a−∼πn(s)
un(s, a+, a−)− β2M 2
8 ,
(31)
and similarly,
g(πh) = E
s∼dpref
E
a+∼πh(s),a−∼πn(s)

−logσ
 
β·u h(s, a+, a−)

≤log 2− β
2 E
s∼dpref
E
a+∼πh(s),a−∼πn(s)
uh(s, a+, a−) + β2M 2
8 ,
(32)
Hence, we have
g(πn)−g(π h)≥ β
2 E
s∼dpref
E
a+∼πh(s),a−∼πn(s)

uh(s, a+, a−)−u n(s, a+, a−)

− β2M 2
4
= β
2 E
s∼dpref
E
a+∼πh(s),a−∼πn(s)

log πh(a+|s)
πh(a−|s) −log πn(a+|s)
πn(a−|s)

− β2M 2
4
= β
2 E
s∼dpref

E
a+∼πh(s)
log πh(a+|s)
πn(a+|s) + E
a−∼πn(s)
log πn(a−|s)
πh(a−|s)

− β2M 2
4 .
(33)
19

## Page 20

By the definition of KL divergence, we have
g(πn)−g(π h) = β
2 E
s∼dpref
h
KL

πh(·|s)
πn(·|s)

+ KL

πn(·|s)
πh(·|s)
i
− β2M 2
4
≥2β E
s∼dpref
h
DTV
 
πh(·|s), πn(·|s)
i2
− β2M 2
4 ,
(34)
where we use Pinsker’s inequality to obtain the bound on DTV
 
πh(·|s), πn(·|s)

from the KL
divergence.
Finally, we apply the inequality E[X 2]≥(E[X]) 2 on X=D TV
 
πh(·|s), πn(·|s)

, so that we have
g(πn)−g(π h)≥2β
h
E
s∼dpref
DTV
 
πh(·|s), πn(·|s)
i2
− β2M 2
4 . (35)
This proves that
E
s∼dpref
DTV
 
πh(·|s), πn(·|s)

≤
s
g(πn)−g(π h)
2β + βM 2
8 . (36)
G Technical Lemmas
Lemma G.1(Expectation Difference via Total Variation).Let P and Q be two probability distri-
butions on a measurable space X , and let f:X →R be any measurable function satisfying the
uniform bound|f(x)| ≤Mfor anyx∈ X. Then
 E
x∼P(·)
f(x)− E
x∼Q(·)
f(x)
 ≤2M·D TV(P, Q),(37)
whereD TV(P, Q) = 1
2 ∥P−Q∥ 1 is the total variation distance.
In addition, when the measurable function g satisfies the bound |g(x1)−g(x 2)| ≤M ′ for any
x1, x2 ∈ X, we have
 E
x∼P(·)
g(x)− E
x∼Q(·)
g(x)
 ≤M ′ ·D TV(P, Q).(38)
Proof.When|f(x)| ≤Mfor anyx, we have
 E
x∼P(·)
f(x)− E
x∼Q(·)
f(x)
 =

X
x
f(x)·(P(x)−Q(x))

≤
X
x
|f(x)| · |P(x)−Q(x)|
≤M·
X
x
|P(x)−Q(x)|
= 2M·D TV(P, Q).
(39)
When |g(x1)−g(x 2)| ≤M ′ for any x1, x2 ∈ X , we set f(x) =g(x)− 1
2 c, where c= sup
x∈X
f(x) +
inf
x∈X
f(x). As we have|f(x)| ≤M ′/2for anyx∈ X, we have
 E
x∼P(·)
g(x)− E
x∼Q(·)
g(x)
 =
 E
x∼P(·)
f(x)− E
x∼Q(·)
f(x)
 ≤M ′ ·D TV(P, Q).(40)
20

## Page 21

Lemma G.2(Performance Gap Between Human Policy and Novice Policy).We denote the Q-function
of human policyπ h asQ ∗(s, a) = E
τ∼P πh
[
∞P
t=0
γtr(st, at)
s0 =s, a 0 =a].
For the human policy πh and the novice policy πn whose value functions are J(π h), J(πn), respec-
tively, we have
J(π h)−J(π n) = 1
1−γ E
s∼dπn

E
ah∼πh(s)
Q∗(s, ah)− E
an∼πn(s)
Q∗(s, an)

. (41)
Proof. We denote the Q-function of novice policy πn as Qn(s, a) = E
τ∼P πn
[
∞P
t=0
γtr(st, at)
s0 =
s, a0 =a].
We denote value functions of πh, πn as V ∗(s) = E
a∼πh(s)
Q∗(s, a) and Vn(s) = E
a∼πn(s)
Qn(s, a),
respectively. And we haveJ(π h) = E
s0∼d0
V ∗(s0), andJ(π n) = E
s0∼d0
Vn(s0).
We define the advantage function ofπ h asA ∗(s, a) =Q ∗(s, a)−V ∗(s).
By the performance difference lemma (Lemma 6.1, [14]), we have
E
s0∼d0

Vn(s0)−V ∗(s0)

= 1
1−γ E
s∼dπn

E
a∼πn(s)
A∗(s, a)

.(42)
This implies that
E
s0∼d0

Vn(s0)−V ∗(s0)

= 1
1−γ E
s∼dπn

E
a∼πn(s)
[Q∗(s, a)−V ∗(s)]

= 1
1−γ E
s∼dπn

−V ∗(s) + E
a∼πn(s)
Q∗(s, a)

= 1
1−γ E
s∼dπn

− E
a∼πh(s)
Q∗(s, a) + E
a∼πn(s)
Q∗(s, a)

.
(43)
Multiplying−1on both sides, we can obtain that
J(π h)−J(π n) = 1
1−γ E
s∼dπn

E
ah∼πh(s)
Q∗(s, ah)− E
an∼πn(s)
Q∗(s, an)

. (44)
H Ethics Statement
Our Predictive Preference Learning from Human Interventions (PPL) delivers a human-friendly,
human-in-the-loop training process that increases automation while minimizing expert effort, ad-
vancing more intelligent AI systems with reduced human burden. All the experiments are conducted
entirely in simulation, ensuring no physical risk to participants. All volunteers provided informed
consent, were compensated above local market rates, and could pause or withdraw from the study
at any time without penalty. Individual sessions lasted less than one hour, with a mandatory rest
period of at least three hours before any subsequent participation. No personal or sensitive data was
collected or shared. We have obtained the IRB approval to conduct this project.
While PPL promises positive social impact by streamlining human-AI collaboration, it may also
encourage overreliance on automated systems or inherit biases present in expert involvement.
21
