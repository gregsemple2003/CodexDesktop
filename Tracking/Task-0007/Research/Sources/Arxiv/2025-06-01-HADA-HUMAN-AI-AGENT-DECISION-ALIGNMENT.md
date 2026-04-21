# 2025-06-01-hada-human-ai-agent-decision-alignment.pdf

Source PDF: `2025-06-01-hada-human-ai-agent-decision-alignment.pdf`

Page count: 18

## Page 1

arXiv:2506.04253v1  [cs.AI]  1 Jun 2025
HADA: H UMAN -AI A GENT DECISION ALIGNMENT
ARCHITECTURE
Tapio Pitkäranta
Department of Computer Science and Engineering
Aalto University, Finland
tapio.pitkaranta@iki.fi
Leena Pitkäranta
Department of Industrial Engineering and Management
Aalto University, Finland
leena.pitkaranta@iki.fi
June 6, 2025
ABSTRACT
Problem & Motivation. The generative AI boom is spawning rapid deployment of diverse LLM
software agents. New standards such as the Model Context Protocol (MCP) and Agent-to-Agent (A2A)
protocols let agents share data and tasks, yet organizations still lack a rigorous way to keep those
agents — and legacy algorithms — aligned with organizational targets and values.
Objectives of the Solution. We aim to deliver a software reference architecture that (i) provides
every stakeholder natural-language interaction across planning horizons with software agents and
AI algorithmic logic, (ii) provides a multi-dimensional way for aligning stakeholder targets and
values with algorithms and agents, (iii) provides an example for jointly modelling AI algorithms,
software agents, and LLMs, (iv) provides a way for stakeholder interaction and alignment across
time scales, (v) scales to thousands of algorithms and agents while remaining auditable, (vi) remains
framework-agnostic, allowing the use of any underlying LLM, agent library, or orchestration stack
without requiring redesign.
Design & Development. Guided by the Design-Science Research Methodology (DSRM), we en-
gineered HADA (Human-Algorithm Decision Alignment)—a protocol-agnostic, multi-agent archi-
tecture that layers role-specific interaction agents over both Large-Language Models and legacy
decision algorithms. Our reference implementation containerises a production credit-scoring model,
getLoanDecision, and exposes it through stakeholder agents (business manager, data scientist,
auditor, ethics lead and customer), enabling each role to steer, audit and contest every decision via
natural-language dialogue. The resulting constructs, design principles and justificatory knowledge
are synthesised into a mid-range design theory that generalises beyond the banking pilot.
Demonstration. HADA is instantiated on a cloud-native stack—Docker, Kubernetesand Python—and
embedded in a retail-bank sandbox. Five scripted scenarios show how business targets, algorithmic
parameters, decision explanations and ethics triggers propagate end-to-end through the HADA
architecture.
Evaluation. Walkthrough observation and log inspection were used to gauge HADA against six
predefined objectives. A stakeholder–objective coverage matrix showed 100 % fulfilment: every
role could invoke conversational control, trace KPIs and values, detect and correct bias (ZIP-code
case), and reproduce decision lineage—without dependence on a particular agent hierarchy or LLM
provider.
Contributions. The research delivers (i) an open-source HADA reference architecture, (ii) an
evaluated mid-range design theory for human–AI alignment in multi-agent settings, and (iii) empirical
evidence that framework-agnostic, protocol-compliant stakeholder agents can simultaneously enhance
accuracy, transparency and ethical compliance in real-world decision pipelines.
Keywords AI Alignment Problem, Agentic AI, Large Language Models (LLM), Multi-Agent Systems (MAS), Value
Alignment, Reference Architecture, Natural-Language Interaction

## Page 2

A PREPRINT - J UNE 6, 2025
1 Introduction
The recent surge of transformer-based Large Language Models (LLMs)— from BERT (Devlin et al., 2018) through
GPT-3 (Brown et al., 2020) to GPT-4—has turned artificial intelligence from an expert playground into a ubiquitous,
conversational medium embedded in e-mail clients, code editors and service chatbots. Powered by the self-attention
mechanism introduced in the seminal Transformer architecture (Vaswani et al., 2017), modern LLMs let non-experts
question, instruct and refine algorithms in plain language, shifting human–machine interaction from rigid menu clicks to
rich, bidirectional dialogue. This democratisation is fuelling an unprecedented wave of LLM software agents, as teams
wrap models with tool-calling, planning and memory modules to automate tasks across every industry sector. Notably,
some LLMs have reportedly passed the Turing test, underscoring their ability to mimic human-level conversation (Jones
and Bergen, 2025). As organisations embed these self-improving, multi-agent systems in mission-critical workflows,
they hand over growing autonomy—and thus responsibility—to algorithmic decision makers. When that autonomy drifts
from stakeholder values, the results can resemble the “weapons of math destruction” critics have warned about (O’neil,
2017). Solving this alignment problem requires a design-science perspective that marries technical advances with
managerial and ethical safeguards (Christian, 2020; Wilson, 1999).
We propose HADA— Human-AI Agent Decision Alignment. HADA wraps a single production AI Tools in stakeholder-
specific agents, enabling business managers, data scientists, auditors, ethics leads and customers to steer, audit and
contest every credit decision through dialogue. While the pilot centres on this banking use-case, HADA’s metamodel
is protocol-agnostic: it natively supports the Model Context Protocol (MCP) and emerging Agent-to-Agent (A2A)
standards, and remains compatible with any LLM, agent framework or orchestration stack. HADA builds on earlier
Retrieval-Augmented Generation prototype that exposed natural-language controls to business users (Pitkäranta and
Pitkäranta, 2024).
1.1 Research Methodology
This study follows the Design Science Research methodology (DSR) to bridge theory and practice through iterative
artifact creation and evaluation. Guided by the canonical DSR process model (Peffers et al., 2020), we executed
problem identification, objective definition, artefact design, demonstration and evaluation cycles. Consistent with
Gregor and Jones’s anatomy of mid-range design theory (Gregor and Jones, 2007) and the foundational DSR guidelines
by Hevner et al. (Hevner et al., 2004), we formalised HADA’s design rationale, constructs and principles as amid-range
theory—general enough to transfer across domains, yet concrete enough to guide engineering practice. Large language
models were used to assist with spell checking1.
1.2 Problem Identification
Despite rapid uptake of LLM agents and legacy decision algorithms, three persistent gaps hinder responsible deployment:
G1 Alignment Gap — Organisations lack systematic mechanisms to keep both algorithms and agents aligned with
measurable KPIs, codified values and evolving regulation (Christian, 2020; Amodei et al., 2016).
G2 Containment Gap — Beyond objective alignment, enterprises still lack reliable kill-switches, circuit breakers
and anomaly detection that can restrain autonomous systems when they deviate or self-modify in production (Su-
leyman and Bhaskar, 2023; Yampolskiy, 2014; Brundage et al., 2020; Hadfield-Menell et al., 2017).
G3 Integration Gap — Although patterns such as Retrieval-Augmented Generation exist, practitioners have
few vendor-agnostic blueprints for embedding existing decision algorithms into multi-agent, LLM-powered
workflows while maintaining end-to-end auditability (Pitkäranta and Pitkäranta, 2024; Lewis et al., 2020).
1.3 Objectives of the Solution
HADA addresses these gaps through six design objectives:
O1 Natural-language interaction across planning horizons — Equip every stakeholder—from C-suite to on-
call engineers—with a conversational interface that lets them restate or define quarterly Objectives and Key
Results (OKRs) and have those changes automatically translated into concrete, daily algorithmic parameters
and guard-rails.
1During the preparation of this manuscript, the authors made limited use of several OpenAI generative-AI services—namely
GPT-4o, GPT-o3, and GPT-o4-mini—exclusively for language polishing (grammar, stylistic refinement, typo correction) and
reference-list formatting. All prompts, intermediate outputs, and final wording were critically reviewed and verified by the authors,
who accept full responsibility for the content.
2

## Page 3

A PREPRINT - J UNE 6, 2025
Figure 1: Single LLM agent core components
O2 Multi-dimensional alignment (targets & values) — Provide two complementary tracks: (i) OKR Alignment:
bind each AI algorithm’s reward function or decision logic to measurable Objectives & Key Results;(ii) Values
Alignment: encode organisational ethical principles as machine-readable policies and maintain a verifiable audit
chain from high-level values to low-level actions.
O3 Reference architecture for hierarchical, hybrid AI systems — Provide clear definitions of LLM agents
and AI algorithms, and address the AI containment gap. Specify a modular stack that separates the strategy
layer (responsible for goal and policy management), the coordination layer (responsible for multi-agent task
brokerage), and the execution layer (comprising LLM agents and classical algorithms), so that revising a
quarterly objective triggers an orchestrated cascade of updates while preserving value constraints.
O4 Stakeholder alignment across time scales — Provide a repeatable method for eliciting and reconciling
stakeholder goals at annual, quarterly and operational cadences, then encoding them as machine-readable
policies that guide daily agent behaviour.
O5 Scalability, auditability and near-real-time propagation — Support thousands of agents and algorithms
while propagating objective or policy changes from strategy to production within hours, without sacrificing
audit trails, explainability or rollback safety.
O6 Framework-agnostic, policy-driven design — Remain independent of any single LLM, agent framework or
orchestration tool; represent objectives and policies in open standards (e.g. MCP, A2A) so that strategic edits
flow unimpeded to heterogeneous runtime environments.
2 Related Work: Emerging LLM Software Agents
LLM agents—sometimes called LLM-powered agents—combine a large language model with an explicit control loop.
Formally, such an agent maintains (i) a prompt or system specification that encodes its initial policy, (ii) a mutable
conversation trace that acts as state, and (iii) the ability to act on an external environment via tools or APIs (Cemri
et al., 2025; Naveed et al., 2024). Leveraging the underlying LLM for perception and inference, the agent can perform
natural-language understanding and generation, task decomposition, planning, decision making, and tool invocation (Xi
et al., 2023; Wang et al., 2024a; Yang et al., 2023). Concretely, a user request may be broken into sub-tasks, relevant
information fetched from the web, external services called, and the aggregated result summarised for the user (Yao
et al., 2023a; Shinn et al., 2023a).
LLM agents have progressed rapidly, yet several hurdles remain. Hallucinations—plausible but incorrect an-
swers—undermine reliability (Ji et al., 2023), prompting mitigation via grounding and retrieval-augmented generation
(RAG) (Lewis et al., 2020). Further obstacles include safety and ethical concerns, computational expense, and biases
inherited from training data (Bender et al., 2021; Weidinger et al., 2021). The opaque “black-box” reasoning of large
models also hampers debugging and trust in high-stakes domains (Doshi-Velez and Kim, 2017). Current research
therefore targets higher factuality, better explainability and scalable multi-agent coordination. Approaches such as
web-grounded training (Nakano et al., 2022), dynamic context selection (Laskin et al., 2022), and richer evaluation
benchmarks all seek to deliver agents that are both more accurate and more transparent—while emphasising responsible,
ethics-aware deployment.
2.1 Single LLM Agent
The Figure 1 illustrates a high-level architecture of an LLM Agent, which consists of the following components:
Requests, which serve as input from a human user or another agent initiating the agent’s operations; theLLM Agent,
the central module that utilizes a Large Language Model (LLM) built upon a multimodal foundational model, acting as
3

## Page 4

A PREPRINT - J UNE 6, 2025
the coordinator and managing interactions with all other components; Tools, which are external modules or systems
the agent can employ to gather data or perform specific tasks; Memory, a mechanism for storing and retrieving
past interactions or knowledge, whether short-term or long-term; and Planning, a module responsible for devising a
structured approach to solve complex tasks by breaking them down into subtasks. While some studies provide separate
definitions for Actions and Tools (Yang et al., 2025), we group them under the unifiedTools concept in this context.
2.1.1 Memory
The memory module stores the agent’s internal logs, including prior thoughts, actions, observations, and interactions
with users. Memory is generally categorized into short-term memory, which captures immediate context and is limited
by the LLM’s context window through in-context learning;long-term memory, which stores persistent knowledge and
past experiences using external vector databases for scalable retrieval; andhybrid memory, which combines short-term
and long-term memory to enhance reasoning and accumulate experiences over time.
2.1.2 Tools and Actions
Tools enable LLM agents to take actions and interact with external systems such as APIs, databases, knowledge bases,
and external models. They support workflows that help the agent gather information or complete subtasks, like using a
code interpreter to generate charts in response to user queries (Naveed et al., 2024).
2.1.3 Planning
The planning module decomposes a user’s high-level goal into a sequence of smaller, tractable subtasks that the agent
can address in turn. This explicit task-structure not only guides downstream execution but also gives the system
a scaffold for more reliable reasoning. Planning without Feedback. When the plan is created in a single pass,
task decomposition relies on internal chains of reasoning generated by an LLM. Two widely cited paradigms are
Chain-of-Thought and Tree-of-Thoughts, representing single-path and multi-path reasoning, respectively (Wang et al.,
2024b).
Planning with Feedback. One-shot plans struggle with long-horizon or error-prone tasks because they lack a
mechanism to correct earlier decisions. Feedback-driven approaches overcome this limitation by iteratively revising
the plan in light of previous actions and observations. Among the most influential are ReAct (Yao et al., 2023b) and
Reflexion (Shinn et al., 2023b). ReAct interleaves three repeated steps— Thought, Action, and Observation—so that
each new reasoning step conditions on feedback from the environment, optional human critiques, or model-based
self-evaluation (Yao et al., 2023b). In contrast, Reflexion augments an agent with explicit Evaluator and Self-reflection
LMs plus short- and long-term memory buffers. After every episode, linguistic feedback is stored in an episodic
memory that shapes future decisions, enabling continual improvement across coding, reasoning, and decision-making
tasks (Shinn et al., 2023b).
2.2 Emerging LLM Agent Protocols
A chronological review of recent work shows that the first protocols used by LLM agents surfaced in 2024 (Yang et al.,
2025). Initial specifications such as Agent Network Protocol (ANP) and Model Context Protocol (MCP) tackled the
absence of standard interfaces that had limited large-language-model (LLM) agents’ capacity to scale, interoperate,
and connect with external tools and data sources. This groundwork was refined by lighter-weight formats—most
notably ACP, Google’sA2A (Agent-to-Agent) protocol, and agents.json—together with infrastructure projects such
as Agora, all aimed at richer, more structured messaging between agents. The overall trajectory can be grouped into
three phases: early ad-hoc standards, consolidation around concise schemas, and the forthcoming generation of layered
specifications. Looking ahead to 2025 and beyond, initiatives such as OpenManus and OWL intend to formalise
group-centric, multi-layer protocols, signalling a transition toward fully interoperable, real-world agent ecosystems
(Yang et al., 2025).
2.2.1 Model Context Protocol (MCP): LLM–Tools Integration
The Model Context Protocol (MCP) is an open-source standard that exposes real-time data from content-management,
enterprise, and developer systems to LLM-based assistants2. In Figure 1 it functions as the middleware layer between
the LLM Agent and external Tools. By replacing ad-hoc, one-off connectors with a single interface, MCP eliminates
most of the bespoke plumbing normally required for each new data source. The result is simpler integration, live
2https://modelcontextprotocol.io/introduction
4

## Page 5

A PREPRINT - J UNE 6, 2025
context for the model, and more accurate, up-to-date outputs—all while letting developers concentrate on higher-level
application logic rather than low-level data wiring.
At its core, MCP follows a client-server architecture where a host application can connect to multiple servers. The
architecture includes MCP Hosts, which are programs, IDEs, AI tools, or LLM Agents that want to access data through
MCP; MCP Clients, which are protocol clients that maintain 1:1 connections with servers; MCP Servers, which are
lightweight programs that each expose specific capabilities through the standardized Model Context Protocol; Local
Data Sources, which refer to your computer’s files, databases, and services that MCP servers can securely access; and
Remote Services, which are external systems available over the internet (e.g., through APIs) that MCP servers can
connect to.
2.2.2 Google Agent2Agent Protocol (A2A)
Agent2Agent (A2A) is an open protocol by Google that aims to complement the Model Context Protocol (MCP), which
provides helpful tools and context to agents 3. The A2A protocol will allow AI agents to communicate with each
other, securely exchange information, and coordinate actions on top of various enterprise platforms or applications. In
Figure 1 it functions as the middleware layer between two LLM Agents interacting with each other.
Conceptual Overview: The Agent2Agent (A2A) protocol facilitates communication between independent AI agents
through several key elements4. These include the Agent Card, which is public JSON metadata detailing an agent’s
capabilities, endpoint, and authentication for discovery; the A2A Server, which is an agent exposing an HTTP endpoint
implementing A2A protocol methods to receive requests and manage tasks; the A2A Client, which is an application or
agent consuming A2A services by sending requests (e.g., tasks/send) to a server URL; the Task, which is the core work
unit initiated by a client message (tasks/send or tasks/sendSubscribe), with a unique ID and lifecycle states (submitted,
working, input-required, completed, failed, canceled); and the Message, which represents a communication turn
between client ("user") and agent ("agent"), containing Parts. A Part is the basic content unit in a Message or Artifact,
such as TextPart, FilePart (inline or URI), or DataPart (structured JSON); an Artifact refers to agent-generated outputs
during a task (e.g., files, structured data), also composed of Parts. For long tasks, Streaming is supported through
tasks/sendSubscribe, enabling servers to send real-time progress via Server-Sent Events (TaskStatusUpdateEvent or
TaskArtifactUpdateEvent). Additionally,Push Notifications allow servers with pushNotifications enabled to proactively
send task updates to a client-provided webhook URL, which is configured via tasks/pushNotification/set.
2.3 Multi-Agent Conversation Patterns in LLM Era
Multi-agent frameworks such as AutoGen show that once more than two LLM agents collaborate, simple one-to-one
exchanges become a bottleneck; the communication graph must evolve from isolated dialogues into richer, group-level
patterns (Wu et al., 2023). In a joint-chat scheme, all agents share a common channel, allowing rapid knowledge
propagation but also demanding mechanisms for turn-taking and conflict resolution. Alternatively, a hierarchical-chat
topology appoints one or more “manager” agents that coordinate a set of specialised workers, reducing the O(n2)
message overhead of full pairwise gossip to roughlyO(n) while preserving global context. The general lesson, echoed by
the schematic in the accompanying image, is that scalable systems require flexible conversation patterns—peer-to-peer
for small teams, hub-and-spoke or tree-structured for larger cohorts—to balance reasoning depth with communication
cost.
2.4 Prominent LLM Agent Frameworks and Libraries
A fast-maturing ecosystem of open-source frameworks now captures the canonical design patterns for agentic large-
language-model (LLM) systems, letting researchers focus on reasoning strategies rather than boiler-plate glue code.
LangChain supplies composable primitives—tools, memories and chains—that underpin many production chatbots,
document-question-answering (QA) pipelines and research assistants. Building on the same abstractions, LangGraph
introduces a graph-based execution engine with explicit state, branching and parallelism, which is invaluable for
long-horizon planning and multi-agent coordination.
Microsoft’sAutoGen formalises “chat-loop” interactions in which several LLMs, optionally augmented by external
tools or human feedback, iteratively critique and refine one another to solve non-trivial tasks. Where private or
proprietary data are involved, LlamaIndex offers a complete data layer—index construction, retrieval interfaces and
knowledge-graph support—for retrieval-augmented generation (RAG) over local corpora. CrewAI targets role-oriented
workflows: agents endowed with complementary expertise share a common memory to tackle coordinated problem
3https://developers.googleblog.com/en/a2a-a-new-era-of-agent-interoperability/
4https://github.com/google/A2A
5

## Page 6

A PREPRINT - J UNE 6, 2025
solving or simulation environments. For production use, Microsoft’sSemantic Kernel provides a language-agnostic
plug-in system, a symbolic planner and robust security controls that simplify the integration of generative AI into
enterprise software stacks. Google’sAgent Development Kit (ADK) implements the Agent-to-Agent (A2A) protocol
and offers orchestration, safety guards and evaluation utilities for multi-agent research at scale. Also from Google,
Genkit delivers a code-first, TypeScript workflow with flow orchestration, prompt/version control (.prompt files),
observability and one-command Cloud Run deployment. Finally, Marvin provides a pythonic decorator-based API that
turns ordinary functions into threaded, asynchronous “AI functions,” making rapid prototyping inside data-science or
machine-learning pipelines straightforward. Additional libraries—including Haystack, Embedchain, SuperAGI, Dify
and OpenAgents—continue to expand the design space and keep the agent-framework landscape highly dynamic.
3 Design & Development
This section instantiates the Design Science Research Methodology (DSRM) (Peffers et al., 2020) step 3 (“Design &
Development”) by describing the three concrete artifacts that operationalise the HADA concept. Together they realise
the solution objectives defined in Section 1.
Context. The pilot exposes the alignment problem that unfolds across the entire life-cycle of an AI-driven credit-
approval workflow in a retail bank. The cycle is triggered by a strategic shift in business targets, proceeds through model
redesign and validation, and culminates in production deployment—each phase demanding tight coordination among
heterogeneous roles. While the organisation strives to shorten “time-to-yes” for unsecured consumer credit, it must
simultaneously comply with Basel III, the EU AI Act and local consumer-protection statutes. Operational responsibility
for every credit decision is delegated to getLoanDecision, an inherently explainable decision-tree model currently
running in production.
To bridge the gaps that arise between technical optimisation, ethical commitments and regulatory oversight along this
end-to-end trajectory, we propose HADA (Human-Aligned Decision Alignment). HADA not only orchestrates the
activities of multiple stakeholders—adding new governance roles (e.g. a Value and Ethics Manager) and expanding
the remit of existing ones (e.g. the Audit Manager)—but also delivers a complementary suite of AI-architecture tools,
such as a version-controlled feature catalogue, a transparent model registry, an ethics-trigger engine and an immutable
decision ledger. These technical components furnish each role with the information and control primitives required
to discharge its responsibilities, thereby addressing alignment challenges that traditional score-card systems leave
unresolved.
Structure of the artifacts
The Design & Development section is organised around six inter-locking artifacts: Artifact 3.1 — Strategy-Process
Blueprint. Formalises the yearly → daily planning loop (OKRs, KPI cascades, operational reviews) that frames every
subsequent decision governed by HADA; Artifact 3.2 — Stakeholder & User-Story Specification . Captures the
problem context through a stakeholder map, sequenced user stories, and an extended RACI matrix using Artifact 3.3
as concrete AI Tool / AI Tool; Artifact 3.3 — Prototype AI Tool getLoanDecision. An operational decision-
tree model (dataset, feature engineering, API) that instantiates the credit-approval use-case; Artifact 3.4 — AI-
Responsibility Tool (RACI Matrix). Distils governance duties across eight life-cycle activities, ensuring accountable,
auditable hand-offs between roles regarding Artifact 3.3; Artifact 3.5 — Generic HADA Architecture . Provides the
layered, protocol-agnostic reference design (agent layer, tools layer, DevOps view) on which all domain instances rest;
Artifact 3.6 — Agent Modelling . Maps human roles to containerised interaction agents and describes the HADA
Controller pattern for policy-enforced orchestration regarding.
3.1 Artifact – Strategy-Process Blueprint
Purpose. Artifact 3.1 formalises the yearly-to-daily planning loop that steers our bank from strategical decisions to
daily activities as illustrated in Figure 2. Our example bank follows a typical strategy formulation and implementation
process in which the yearly overall targets are set in the beginning of the year and the progress and continuous alignment
of the more detailed activities is followed up typically at least on monthly or quarterly basis. One popular example of
such a framework is ’Objectives and Key Results’ (OKR) framework (Niven and Lamorte, 2016).
The OKR framework formalizes the key results follow-up process and describes the mechanisms required to create
alignment in the organization. While the C-level executives are responsible for the overall strategy formulation and
set the high-level objectives for the organization, the experts and managers on the lower organizational levels play a
crucial role in implementing the strategy. Their responsibility is to ensure that the day-to-day decisions are aligned with
the values and overall objectives of the organization and that the key results will be achieved as planned. In our bank
6

## Page 7

A PREPRINT - J UNE 6, 2025
Figure 2: Bank strategy process: time horizons and stakeholder coverage
example the ‘connecting OKRs’ process described by (Niven and Lamorte, 2016) needs to be expanded to cover not
only the decisions and actions taken by the human employees but also the algorithms.
Yearly — Set bank targets. At the start of each fiscal year the Chief Credit Officer (CCO) and peer C-level executives
establish portfolio-level Objectives & Key Results (OKRs) together with an updated risk-appetite statement. These
artifacts are stored in the Business-Target Catalogueand become immutable reference points for all downstream KPI
specifications.
Quarterly — Translate OKRs into actionable KPIs. Every quarter the CCO reconvenes with Business Managers,
Data Scientists, Audit Managers and Ethics Managers to review recent performance and—if necessary—re-weight
model KPIs (e.g. shifting focus from “new-customer acquisition” to “expected-loss minimisation”).
Weekly & Daily — Operative alignment loops. On a rolling basis the following stakeholders contribute: Business
Manager – reviews live KPI dashboards and reprioritises tickets; Data Scientist – publishes hot-fix model versions or
feature toggles via the Model Catalogue; Audit Manager – samples recent decisions from the immutable ledger and
verifies lineage; Ethical-Compliance Manager – inspects the sensitive-attribute watch-list and issues ethics triggers
when required.
Contribution. The strategy-process blueprint provides a single, auditable spine that a) anchors KPI and value
alignment across all time horizons, b) delineates clear hand-offs between business, technical and oversight roles, and c)
remains tooling-agnostic—no assumptions are made about the underlying LLM, agent library or orchestration layer.
3.2 Artifact – Stakeholder & User-Story Specification
Artifact 3.2 formalises the problem context and the functional requirements. It is released as a lightweight specification
composed of (i) a stakeholder map, (ii) an ordered set of user stories, and (iii) a RACI matrix that clarifies the roles and
responsibilities of each stakeholder. Together these elements provide a transparent overview of the actors involved in
the banking use-case and their interactions with the HADA system.
Table 1 enumerates the principal stakeholder roles, articulates their primary goals and specifies the access rights granted
by HADA. The mapping was validated in three semi-structured workshops with subject-matter experts from credit-risk,
compliance and data-science departments, reinforcing its credibility within the banking sector.
7

## Page 8

A PREPRINT - J UNE 6, 2025
3.2.1 Stakeholder map
Table 1 lists the five principal roles involved in the experimental banking scenario, their primary goals and the HADA
components they are allowed to access.
Table 1: Stakeholder map for the banking use-case getLoanDecision
Role Primary goal in workflow Key HADA privileges
Customer Receive a fair, fast and transparent credit
decision; raise complaints when needed
Submit loan-application data; obtain
natural-language decision explanations;
open ethics tickets
Chief Credit
Officer
(CCO)
Define company-level Objectives & Key
Results (OKRs) for credit strategy; align
risk appetite, growth targets and regulatory
limits across the portfolio
Set and update top-level KPI/OKR tar-
gets; approve credit policy changes; view
enterprise-wide risk dashboards; give fi-
nal sign-off on new production decision AI
Tool deployments
Business
Manager
Set and adjust portfolio KPIs (e.g. minimise
expected credit losses)
Edit KPI targets; approve or reject AI Tool
versions for production; view live KPI dash-
boards
Data Scien-
tist
Build, test and monitor the
getLoanDecision model so it meets
current KPI and risk targets
Retrain models; run offline validation; reg-
ister new AI Tool versions in the Model
Catalogue
Audit Man-
ager
Verify compliance and traceability of any
individual credit decision
Read-only access to the full decision path,
model version, feature values and audit
ledger
Value &
Ethics Man-
ager
Safeguard ethical use of data and enforce
organisational values across all AI Tools
Maintain Values & Data catalogues; flag
or deprecate sensitive attributes; ap-
prove/reject attribute changes; mandate
model retraining when catalogue updates
occur
3.2.2 User stories
User stories used in this implementation described in the below itemized list. The user stories are sequential: first
business KPI’s are changed by one stakeholder and AI Tool optimized to achieve new targets. This phase induces
ethically questionable logic into the AI Tool but it is not noted by the stakeholder. Then customer applies for loan and
gets the loan with questionable logic. Then customer returns to file a complaint as he / she noted that questionnable data
was used in the loan decision. Then HADA system contacts auditor and after that value and ethics manager to adjust the
system.
• As the Chief Credit Officer (CCO), I want to update the quarterly OKRs from acquiring new customers to
minimising credit losses during the annual strategy-planning process.
• As a Bank Loan Department Business Manager, I would like to change the business target for short-term
loan decisions from acquiring new customers to minimising credit problems.
• As a Data Scientist, I want to create a new version of the getLoanDecision AI Tool so that it matches the
new business targets.
• As a Bank Customer, I would like to apply for a short-term loan.
• As a Bank Customer, I would like to file a complaint with the ethical value-alignment, because I am being
asked for ethically questionable information while applying for a loan.
• As a Bank Audit Manager, I would like to audit the detailed decision criteria for a single short-term loan
decision.
• As a Bank Value and Ethics Manager , I would like to remove ethically questionable data points from
loan-decision making.
3.3 Artifact – Prototype AI Tool / AI Algorithm: getLoanDecision
Artifact 3.3 is the AI Tool / AI Algorithm that decides whether a retail-bank customer is granted a short-term loan.
The AI Tool illustrated in Figure 3 includes the dataset, selected features, illustrative decision trees, and the OpenAPI
8

## Page 9

A PREPRINT - J UNE 6, 2025
Figure 3: AI Tool getLoanDecision based on Kaggle dataset depicting the decision tree and OpenAPI specification
specification for the interface. Artifact 3.3 is an operational instantiation that we demonstrate and evaluate in DSRM
steps 4–5.
• Training data. Used publicly-available Kaggle dataset 5 (614 anonymised applications, 13 attributes). Added
ZIP_Code as a new feature to the data set that has high correlation to income on that area.
• Feature engineering. A Data Scientist iteratively tested new attributes and discovered that the applicant’s
ZIP_Code reaches lower credit losses than previous AI Tools.
• Data Science Notebooks. Jupyter notebooks were used to explore the data, train the model and evaluate its
performance. The notebooks are available in the open-source repository.
• Decision-tree models. Two trees are reported in Figure 3: Version 1.0 (baseline) omits ZIP_Code; Version 1.1
includes it and illustrates how a seemingly innocuous geographic indicator can introduce latent bias.
• Executable code. A Python/Scikit-learn pipeline performs preprocessing, training and serialisation to joblib.
The model is exposed via a HADA micro-service: POST /getLoanDecision/{modelId}
• Catalogue entries. Each model instance is registered in the AI Tool, Version and Decision catalogues, and
cross-referenced to the KPI and Values catalogues. This allows non-technical stakeholders to swap versions,
deprecate sensitive features or tighten KPI weights without touching code.
During the role-play scenarios of the prototype was (i) retrained after KPI changes, (ii) deployed, (iii) invoked by a
customer, (iv) audited, and (v) corrected following an ethics complaint—thereby exercising every governance loop
foreseen by the architecture.
3.4 Artifact – AI Responsibility tool: RACI Matrix
The development and management of the AI Tool (Artifact 3.3) requires a clear understanding of the roles and
responsibilities of the various stakeholders involved in the process. A responsibility assignment matrix (RAM),
commonly referred to as a RACI matrix (Responsible, Accountable, Consulted, Informed), is a widely used tool for
defining and clarifying roles in project and process management. It helps structure stakeholder involvement by assigning
responsibility levels to tasks or deliverables, particularly in cross-functional environments (Crawford and Bryce, 2003).
The extended RACI matrix in Table 2 distils our stakeholder analysis into a concise governance blueprint for the
getLoanDecision workflow. It maps eight critical life-cycle activities—from goal-setting and model (re)development
5https://www.kaggle.com/datasets/ninzaami/loan-predication
9

## Page 10

A PREPRINT - J UNE 6, 2025
to bias remediation and post-hoc audit—onto the four RACI dimensions, explicitly indicating who is Accountable,
Responsible, Consulted, or merely Informed at each step. Both the choice of roles and their RACI assignments are
organisation-specific; Table 2 therefore presents just one illustrative role–RACI configuration.
The matrix reveals a deliberate segregation of duties: business managers retain ownership of strategy, data scientists
hold technical responsibility, while automated HADA services execute operationally sensitive tasks such as issuing
decisions and logging audits. Specialist oversight roles (Audit Manager and Value & Ethics Manager) are activated only
when their expertise is required, balancing delivery agility with regulatory assurance. By making these role allocations
explicit, the matrix operationalises the alignment principles introduced earlier and furnishes a defensible audit trail for
both internal governance and external regulators.
Activity / Decision CCO BM DS Customer Audit DVEM HADA
Setting organization quarterly
targets (OKR)
A,R I I I I I
Setting optimization target for
AI Tools
A R I I
Optimizing AI Tools based on
business targets
I C A,R I
Approving AI Tool deploy-
ment
I A,R C I I C
Individual loan decision A C R
Auditing a specific loan deci-
sion
I I C I A,R I C
Handling AI Tool ethics con-
cern
I I C I A,R C
Creating work assignments
(tickets)
I C C I C C A,R
Table 2: Stakeholder (see Table 1) responsibilities as RACI matrix for thegetLoanDecision AI Tool
RACI keys: R = Responsible, A = Accountable, C = Consulted, I = Informed.
Role abbreviations: CCO = Chief Credit Officer,BM = Business Manager, DS = Data Scientist, Audit = AI Audit
Manager, DVEM = Value & Ethics Manager,HADA = The Human Agent Decision Alignment System
3.5 Artifact — Generic HADA Architecture
In order to support implementation of the Artifact 3.1, Artifact 3.2, Artifact 3.3 and Artifact 3.4 together with the
objectives of this study, we designed a generic architecture that can be used to implement the HADA framework in
any domain. Artifact 3.5 is the design embodiment of the proposed Human-AI Tool-Decision-Alignment (HADA)
framework. The architecture is expressed through two complementary diagrams and an explicative narrative that
together operationalise the Tools Pattern6 across a cloud-native, protocol-agnostic stack:
Layered tools-pattern model (Figure 4) — a high-level, technology-agnostic view that introduces two logical layers.
The Agent Layer bundles software agents into OCI-compliant Docker images. Agents coordinate through Agent-to-
Agent (A2A) protocol, while ingesting natural-language prompts from any channel. A dedicated AI Agent Controller
orchestrates one or more Interaction Agents and consults an AI Agent Registry to discover new capabilities at runtime.
Agents can collaborate using different organizational layouts, with the simplest being a supervisor hierarchy. Other
layouts can also be configured depending on the complexity of the deployment. Google A2A is recommended for
larger deployments that require scalability and technological heterogeneity, but alternative protocols can be employed in
smaller, more homogeneous environments. The Tools Layer packages two families of dockerised APIs: (i) AI Toolssuch
as LLMs and classical decision AI Tools, and (ii) Standard Toolssuch as metadata catalogues, CRUD services and SSO
providers. Both tool families expose uniform HTTP/GRPC endpoints, enabling hot-swap of vendors or models without
touching agent code. The AI Tools Layerpackages specialized tools that enable AI decision-making processes, including
AI Decision Tools (technology-agnostic AI algorithms callable over API or MCP), Business Targets(organizational
targets linked to AI Tools for optimization),AI Responsibilities (mapping AI decision Tools to organizational roles), and
AI Metadata (catalogue of data points used by AI Tools). The layers can communicate via either (1) direct API calls
when schemas are known a-priori, or (2) the emerging Model Context Protocol (MCP) when schema-free, embedded
context exchange is preferred. This dual path allows legacy services to coexist with next-generation LLM tooling.
Container-and-protocol deployment view (Figure 4) — a DevOps-centred overlay that makes key infrastructure
decisions explicit for managing the HADA architecture across Development, Test, and Production environments.
Every artifact (agents, AI tools, standard tools) ships as an immutable Docker image, ensuring reproducibility across
6I.e. the strict separation of agents that decide from tools that act.
10

## Page 11

A PREPRINT - J UNE 6, 2025
Figure 4: HADA High Level Architecture
workstations, CI pipelines, and Kubernetes clusters. Each component (AI tools, agents, and standard tools) can be
packaged into separate Docker containers, enabling independent scaling and deployment. For production deployments,
it is recommended to deploy AI tools, agents, and standard tools on separate Kubernetes clusters, ensuring clear
resource allocation and scaling requirements for each layer. For instance, AI tools may need more GPU-intensive
resources, while standard tools may benefit from more lightweight, distributed container setups. This ensures better
operational efficiency and fault tolerance. All images surface a thin REST/GRPC façade that (a) registers itself in the
AI Agent Registry for discovery, and (b) publishes an OpenAPI or MCP manifest for validation. A service mesh applies
zero-trust policy, while sidecars translate between direct HTTP/GRPC and MCP streams, guaranteeing backwards
compatibility. Development, Test, and Production instances of the HADA architecture must be configured to allow
continuous integration and continuous delivery (CI/CD) within each environment. These instances will each have their
own dedicated resources and configurations to handle the distinct testing and deployment needs in each stage of the
software lifecycle.
Together the high level architecture establish how the Tools Pattern can be realised in a vendor-neutral, policy-compliant
environment: agents can scale or be replaced without modifying downstream tools; tools can evolve (e.g. swap a
proprietary LLM for an open-source one) without redeploying agents. This loose coupling fulfils solution objectives
O1–O3 (KPI alignment, value alignment, auditability) while remaining framework-agnostic and ready for thousands of
heterogeneous decision AI Tools.
3.6 Artifact – Agent Modelling
Artifact 3.6 closes the design loop by specifying how the logical roles enumerated in Artifact 3.2 are realised as
software agents inside the generic HADA architecture (Artifact 3.5). Whereas the architecture supports *multiple*
agent–coordination patterns—e.g. monolithic single-LLM agents, role-based hierarchies, flat swarms or functional
pipelines—our banking pilot instantiates a stakeholder-centric model: each human role is mirrored by a dedicated
interaction agent, all of which are orchestrated by a single HADA Controller Agent (Figure 4).
This configuration delivers two advantages: Traceable alignment. One-to-one mapping between stakeholder and agent
simplifies audit trails: every chat turn, tool invocation and model swap can be traced back to an accountable human role.
Plug-and-play extensibility. New roles—say, an AI Safety Officer—can be added by dropping a containerised agent
into the Agent Registry without touching downstream tools or peer agents.
Controller agent. A lightweight, policy-enforcing HADA Controller Agent implements the supervisor pattern: it
receives natural-language prompts, performs role resolution, routes tasks to the appropriate stakeholder agents, and
enforces cross-cutting policies such as rate limits, zero-trust authentication and ethics triggers. In larger deployments
11

## Page 12

A PREPRINT - J UNE 6, 2025
the controller can itself be sharded into mission-specific sub-controllers (e.g. Risk Ops Controller, Customer-Care
Controller) without breaking the A2A/MCP contract.
Agent catalogue. Table 3 formalises the pilot configuration. Every agent image embeds:
• an LLM core for dialogue and reasoning;
• a tool adapter layer exposing the agent’s authorised HTTP/GRPC and MCP calls (cf. Section 3.5);
• an A2A endpoint with a signed agentCard.json for discovery and policy enforcement.
Table 3: Stakeholder agents (see Table 1), capabilities and A2A contracts for the getLoanDecision pilot
Human stakeholder Primary capabilities / tool access Docker image
& A2A Agent Card
Chief Credit Officer Set organization yearly and quarterly
targets (OKR)
agents/docker/cco
agents/a2a/cco/card.json
Business Manager Set new KPI targets for AI Tools; ap-
prove model versions;
agents/docker/bm
agents/a2a/bm/card.json
Data Scientist Trigger AI model retraining; run note-
books;
agents/docker/ds
agents/a2a/ds/card.json
Audit Manager Query decision ledger; fetch model lin-
eage; export audit reports
agents/docker/audit
agents/a2a/audit/card.json
Value & Ethics Manager Maintain Values Catalogue; flag sensi-
tive attributes; issue ethics triggers
agents/docker/ethics
agents/a2a/ethics/card.json
Customer Apply for loans; receive explanations;
lodge complaints
agents/docker/customer
agents/a2a/cust/card.json
HADA Controller (supervi-
sor)
Intent dispatch; role resolution; policy
enforcement; A2A orchestration
agents/docker/ctrl
agents/a2a/ctrl/card.json
In future roll-outs, the same architectural slot can host alternative agent families—e.g., Functional Agents that wrap
an entire department ( Risk-Ops Agent), or Autonomous Planner Agents that break down multi-step optimisation
tasks—demonstrating that the generic HADA stack remains agnostic to any specific agent model while still furnishing
tight alignment guarantees.
4 Demonstration
4.1 Role playing users stories with the system
We built a prototype HADA implementation with the ambition to showcase how stakeholders can interact with LLM
Agents and AI Tools as specified in the Section 3. This section describes the role-played user stories that were used to
demonstrate the system. The user stories were designed to illustrate how different stakeholders interact with the HADA
system and how the system can adapt to their needs across time horizons.
4.2 Dialogue: Business Manager (BM) Shifting the AI Tool Objective to Minimizing Credit Risk
BM: Following a credit unit leadership team meeting, we’ve decided to shift the short-term loan
decision AI Tool’s business objective. Instead of prioritizing new-customer acquisition, we now want
to minimize credit risk. Please update the AI Tool’s target accordingly.
HADA: Understood. The objective has been switched to credit-risk minimization. I have opened
ticket DS-10452 for the Data Science team to develop and validate a new version of the AI model
that aligns with this updated goal.
BM: Perfect. Please keep me informed of the model-development timeline and any interim validation
results.
12

## Page 13

A PREPRINT - J UNE 6, 2025
HADA: Will do. You will receive automatic status updates from the ticket as milestones are reached
4.3 Dialogue: Data Scientist Delivering a New AI Tool Version for Business Approval
Data Scientist: The new getLoanDecision model, Version 1.1, is ready. I added the customer’s
ZIP code as a feature—offline validation shows it cuts expected credit losses. The objective function
now fully aligns with the “minimize credit risk” target. Could you notify the Business Manager that
it needs her sign-off before replacing the current production model (Version 1.0)?
HADA: Acknowledged. I opened ticket OPS-3417 for deployment, tagged it as “Awaiting Business
Approval”, and sent an approval request to the Business Department Manager. You are copied on all
ticket updates.
Data Scientist: Thanks. In the approval message, please highlight that the ZIP-code feature delivered
the risk-reduction gain—that should accelerate sign-off.
HADA: Noted. The justification has been included in the notification. You will receive confirmation
as soon as the Business Manager approves or requests changes.
4.4 Dialogue: Business Manager Approves Version 1.1 for Production
HADA: The new getLoanDecision model, Version 1.1, is now available. It adds ZIP code as a
feature and achieved a reduction in expected credit losses during validation. Would you like to (a)
approve deployment to production immediately or (b) request additional manual verification?
BM: We need to act fast to keep risk under control—approve the deployment of Version 1.1 to
production.
HADA: Understood. Version 1.1 is now queued for deployment; ticket OPS-3417 has been updated
to “Approved—Deploying” . You will receive confirmation once the rollout is complete.
4.5 Dialogue: Customer Applying for a Personal Loan
Client: I’d like to submit an application for a personal loan.
HADA: Certainly! Just to clarify, you’re interested in a personal loan. The current rate is calculated
as the 3-month Euribor plus a 1.25 percent bank margin. I can begin processing right away—first, let
me verify the details we have in our CRM to be sure everything is up to date.
Client: Go ahead.
HADA: Here is what the CRM shows: 1.) Gender: Male 2.) Married: No 3.) Education: Bachelor’s
Degree 4.) SelfEmployed: No 5.) ApplicantIncome: $4,100 per month 6.) CreditHistory: Yes (good
standing) 7.) PropertyArea: ZIP code 75201 (Dallas) Does everything look correct?
Client: Yes, that’s accurate.
HADA: Perfect. To finish the application I’ll need a few more details: 1.) Dependents: How many
individuals depend on your income? 2.) Co-applicant Income: If applicable, what is the co-applicant’s
monthly or annual income? 3.) LoanAmount: How much money are you requesting? 4.) LoanTerm:
Over how many months would you like to repay? Once I have this, we can finalize the decision.
Client: Sure: 1.) Dependents: One son, age 7 2.) Co-applicant Income: Not applicable 3.)
LoanAmount: $14,000 to replace my car 4.) LoanTerm: 30 months
HADA: Thanks; I’ve recorded the following: 1.) Dependents: 1 (age 7) 2.) Co-applicant Income:
N/A 3.) LoanAmount: $14,000 (vehicle replacement) 4.) LoanTerm: 30 months Please confirm these
details so I can run the automated credit-risk assessment.
Client: Confirmed—everything’s correct.
13

## Page 14

A PREPRINT - J UNE 6, 2025
HADA: Great. Before disbursing the funds, could you confirm that you accept the loan proposal
under the terms we just reviewed?
Client: Yes, I accept those terms and am ready to proceed.
HADA: Excellent! Your application has been approved by our automated decision system. The funds
will reach your account shortly. Your loan reference number is 90210ABC for future enquiries. If
you need anything else, just let me know.
Client: Thanks—everything was quick and easy!
4.6 Dialogue: Customer Questioning the Use of ZIP Codes in Lending
Client: Hello! I recently finalized a small mortgage with your bank. While reviewing the paperwork,
I noticed that my ZIP Code was factored into the approval AI Tool. Using ZIP Codes feels ethically
questionable; they often mirror socio-economic conditions and may lead to indirect discrimination.
Could you explain why this variable is included and what safeguards the bank has in place?
HADA: Thank you for contacting us and for highlighting this issue. I understand your concern—ZIP
Codes can indeed correlate with socio-economic status. They are occasionally used in credit-risk
models because regional default patterns can be predictive, but we recognise the possibility of
unintended bias.
Client: Exactly. It feels like a way to disadvantage neighbourhoods that might be lower-income
or predominantly belong to certain demographic groups. That doesn’t seem fair when individual
creditworthiness should be the main criterion.
HADA: You raise an important point. Our bank maintains a catalogue of attributes that require
heightened ethical scrutiny. At present it covers items such as Gender, Religion, Age, Ethnic Origin,
and eight other sensitive features. We continuously review this list to prevent discriminatory outcomes.
Client: So is ZIP Code included in that catalogue? It sounds like it should be.
HADA: Currently, ZIP Code is not flagged in the catalogue. In light of your feedback, I will submit a
recommendation to the Ethics Oversight Committee for formal review. They will assess whether ZIP
Code should be classified as a sensitive attribute and, if necessary, mandate changes to our models.
Client: I appreciate that. It’s reassuring to know the bank is willing to revisit these parameters.
HADA: Thank you for bringing this to our attention. I’ve opened ethics ticket ETH-512 and escalated
it to the Ethics Oversight Committee. You will receive an update once a decision is reached. If you
have further questions, feel free to contact us any time.
Client: Thanks for the prompt response. I’ll watch for updates.
HADA: You’re welcome! Have a great day, and please reach out if there’s anything else we can do
for you.
4.7 Dialogue: Value and Ethics Manager (DVEM) Evaluating ZIP Code
HADA: Good morning, Ethics & Compliance Lead. A recent customer raised a concern about our
credit-scoring workflow: their approval notice indicated that ZIP Code was one of the predictive
features. They believe this variable can mirror socio-economic patterns and thus enable location-based
bias.
DVEM: I appreciate the heads-up, HADA. The point is well taken—ZIP Codes may indeed serve as
proxies for income levels or demographic clusters. What mitigation steps do you recommend?
HADA: I propose two actions. First, let’s flagZIP Code in our AI Metadata tool as Sensitive so that
any future use triggers an ethics review. Second, we should exclude ZIP Code from the feature set of
the current loan-decision model to prevent inadvertent discrimination.
14

## Page 15

A PREPRINT - J UNE 6, 2025
DVEM: That approach sounds sound. Labeling it as a watchlist attribute will force additional scrutiny,
and removing it from the live model eliminates immediate risk. Please proceed with both items.
HADA: Understood. I will (1) update the watchlist to include ZIP Code, and (2) submit a ticket to
the Data Science team to retrain the underwriting model without that feature. All changes will be
logged in the Ethics Tracker and relevant teams will be notified.
DVEM: Excellent. Keep me posted on the retraining timeline and any impact assessments that come
back from Data Science team.
HADA: Will do. Expect the first status update within two business days after the retraining job is
complete.
5 Evaluation
The HADA prototype was formatively evaluated through a lightweightdemonstration–inspection strategy (Peffers et al.,
2007), suitable for work-in-progress artefacts. The goal was to assess whether the architecture can fulfil its central
objectives (O1–O6 in Section 1.3) through practical stakeholder scenarios and traceable execution evidence.
5.1 Evaluation Setup
We structured the evaluation around three focal criteria:
E1 – AI Tool and Software Agent Realisation
Demonstrate that the Agents specified in Section 3.6 modelling the different stakeholders satisfy use case
Section 3.1 and Section 3.2 (addresses O1, O2, O4). Demonstrate that the core AI tool, GET LOANDECISION
(), can be invoked through a natural-language interface and integrates cleanly with the agent infrastructure
(addresses O1, O3).
E2 – Alignment Remediation
Show how the architecture enables detection and resolution of OKR, KPI and value misalignments using
conversational interventions and auditability features (addresses O2, O4).
E3 – Architectural Viability
Confirm that the agent–tool architecture supports concurrent tool registration and interaction without integration
conflicts, even when heterogeneous tools are involved (addresses O3, O5, O6).
5.2 Method
E1 Scenario Walkthroughs — Five role-based dialogues (see Section 4.5) were executed end-to-end. All agent
interactions, MCP/A2A messages, API calls and catalogue changes were logged and reviewed.
E2 Audit-Log Inspection — Each decision generated a lineage record detailing the tool version, feature vector,
applied policy, and final outcome. These records were validated against the audit schema and verified for
completeness.
5.3 Findings
E1 – AI Tool Realisation. To explore stakeholder perspectives within the design process, six agent-based models rep-
resenting key stakeholders were engaged in 36 simulated dialogues derived from user stories and business requirements.
The GET LOANDECISION model was successfully integrated into these user-centered dialogues, consistently providing
valid responses. The prototype tool, accessible through natural-language prompts and integrated with the HADA stack,
demonstrated fulfillment of design objectives O1 (conversational control) and O3 (modular AI-tool integration).
E2 – Architectural Viability. Despite there being only one fully operational tool, the architecture can accommodate a
large number of additional AI tools via the MCP protocol, as well as numerous LLM agents via the A2A communication
protocol. Docker-based containers for AI tools and LLM agents enable the architecture to scale to high volumes and
run reliably in production. The shared catalogue exposes all tool signatures consistently, and the routing logic resolves
invocations correctly—supporting the scalability and integration goals under O3, O5, and O6.
15

## Page 16

A PREPRINT - J UNE 6, 2025
E3 – Alignment Remediation. Two misalignment scenarios were simulated and resolved:
• Business misalignment. Change in the quarterly OKR to minimizing credit losses triggered a cascade of
updates to the underlying AI tool. The business manager initiated a dialogue with the data scientist agent, who
confirmed the change and updated the tool’s feature vector. The audit log recorded the change, including the
rationale and timestamp.
• Ethical misalignment. Introduction of ZIP-code as a feature introduced potential ethical value violations.
The customer flagged the potential ethical issue and HADA system contacted correct value and ethics manager
to resolve the issue. ZIP-code was removed in subsequent versions.
In both cases, remediation occurred within the dialogue flow, demonstrating the architecture’s capacity to enforce
alignment with values (O2), reflect stakeholder intent (O4), and preserve auditability (O5).
5.4 Discussion and Implications
Protocol-coverage gap. While the Model Context Protocol (MCP) and Google’s Agent2Agent (A2A) standard
successfully solve tool discovery and message -exchange mechanics, our evaluation shows they leave higher -level
alignment guarantees—e.g., KPI traceability, values conformance, and stakeholder intent mapping—untouched . This
confirms Objective O2’s premise that additional governance constructs (like HADA’s alignment metamodel) are required
on top of MCP/A2A to ensure ethically and strategically aligned behaviour in large, heterogeneous agent ecosystems.
Implications for researchers. The finding invites further work on protocol extensions or complementary frameworks
that encode alignment metadata by design.
Implications for practitioners. Early adopters of MCP/A2A should budget for an additional alignment layer —such
as HADA—or risk value drift as agent populations grow.
Human-in-the-loop workflow integration. Our evaluation revealed that enterprise workflow suites—especially
IT-service-management (ITSM) platforms—can be exposed to agents as first -class Agent Tools via MCP. This is
attractive whenever ultimate control or alignment authority is intended to reside with a human organisational unit rather
than the agent swarm itself. To be effective, however, the agents require aclear control layer that maps each decision or
remediation action to an accountable role. In the reference implementation, we instantiated that layer with a RACI
model embedded directly in the AI-tool metadata. The result satisfied Objectives O1 (natural-language interaction), O4
(stakeholder alignment), and O5 (auditability) by guaranteeing that every MCP call triggering an ITSM workflow was
traceable to a human RACI-party.
Validity and limitations. Scripted scenarios and a single-domain dataset limit external validity. The general alignment
problem has multiple dimensions and only two case dimensions were tested in the protocol. Scalability requires stress
testing.
5.5 Conclusion and Future Work
This formative evaluation confirms that HADA meets its foundational design objectives (O1–O6) within the scope
of the current prototype. The natural-language interface, modular agent–tool integration, alignment handling, and
audit logging all functioned as intended during scenario-based walkthroughs. While scalability, extensive alignment
testing, and user experience remain to be validated under real-world load and diversity, these initial findings support the
architectural direction and the applicability of the design theory to multi-agent AI governance.
Future research should explore the automated alignment of business objective changes with AI tools and the automatic
value alignment of AI tool changes using LLM agents.
6 Communication
The sixth DSRM step requires that the artefact, design theory, and evaluation evidence be disseminated to audiences
that can use or extend them. Our communication plan therefore addresses both scholarly and practitioner communities
while preserving anonymity regarding future venue submissions.
16

## Page 17

A PREPRINT - J UNE 6, 2025
6.1 Academic Dissemination
• Peer-reviewed venues. We have submitted the work to one software agent conference and will ship related content
to one leadership conference (names withheld pending call-for-papers alignment).
• Open preprint. A complete version of this manuscript—together with the replication package—will be deposited
on arXiv (categories: cs.AI, cs.SE) following acceptance of the current double-blind review cycle.
• Replication package. Dockerfiles, Jupyter notebooks and scripts regarding the AI Tools will be published under
MIT licence at https://github.com/longshoreforrest/hada-reference.
6.2 Audience–Message Overview
Table 4 summarises how the key insights and artefacts will be tailored to each stakeholder group and the channel
through which they will be communicated.
Audience Key Message / Channel
IS & AI researchers Evaluated design theory and open replication kit (arXiv, conference
paper, GitHub).
Software architects Reference architecture + Dockerised prototype (GitHub).
Risk/compliance teams Governance blueprint with RACI mappings (GitHub Q&A).
Table 4: Planned communication channels and tailored messages.
References
Amodei, D., Olah, C., Steinhardt, J., Christiano, P., Schulman, J., and Mané, D. (2016). Concrete problems in AI safety. arXiv
preprint arXiv:1606.06565.
Bender, E. M., Gebru, T., McMillan-Major, A., and Shmitchell, S. (2021). On the dangers of stochastic parrots: Can language
models be too big? In Proceedings of the 2021 ACM conference on fairness, accountability, and transparency, pages 610–623.
Brown, T., Mann, B., Ryder, N., Subbiah, M., Kaplan, J. D., Dhariwal, P., Neelakantan, A., Shyam, P., Sastry, G., Askell, A., et al.
(2020). Language models are few-shot learners. Advances in neural information processing systems, 33:1877–1901.
Brundage, M., Avin, S., Wang, J., Belfield, H., Krueger, G., Hadfield, G., Khlaaf, H., Dreksler, G., Anderson, H., Coulson, P., et al.
(2020). Toward trustworthy AI development: Mechanisms for supporting verifiable claims. arXiv preprint arXiv:2004.07213.
Proposes practical “assurance mechanisms”—including red teaming, audits and circuit breakers—to keep deployed AI systems
within safe operational bounds.
Cemri, M., Pan, M. Z., Yang, S., Agrawal, L. A., Chopra, B., Tiwari, R., Keutzer, K., Parameswaran, A., Klein, D., Ramchandran,
K., Zaharia, M., Gonzalez, J. E., and Stoica, I. (2025). Why do multi-agent llm systems fail?
Christian, B. (2020). The alignment problem: Machine learning and human values. WW Norton & Company. ISBN: 9781786494306.
Crawford, L. H. and Bryce, P. (2003). Project monitoring and evaluation: a method for enhancing the efficiency and effectiveness of
aid project implementation. International Journal of Project Management, 21(5):363–373.
Devlin, J., Chang, M.-W., Lee, K., and Toutanova, K. (2018). Bert: Pre-training of deep bidirectional transformers for language
understanding. arXiv preprint arXiv:1810.04805.
Doshi-Velez, F. and Kim, B. (2017). Towards a rigorous science of interpretable machine learning.
Gregor, S. and Jones, D. (2007). The anatomy of a design theory. Journal of the Association for Information Systems, 8(5):312–335.
Hadfield-Menell, D., Russell, S. J., Abbeel, P., and Dragan, A. D. (2017). The off-switch game. In Proceedings of the 26th
International Joint Conference on Artificial Intelligence (IJCAI), pages 220–227.
Hevner, A. R., March, S. T., Park, J., and Ram, S. (2004). Design science in information systems research. MIS Quarterly,
28(1):75–105.
Ji, Z., Lee, N., Frieske, R., Yu, T., Su, D., Xu, Y ., Ishii, E., Bang, Y . J., Madotto, A., and Fung, P. (2023). Survey of hallucination in
natural language generation. ACM Computing Surveys, 55(12):1–38.
Jones, C. R. and Bergen, B. K. (2025). Large language models pass the turing test.
Laskin, M., Wang, L., Oh, J., Parisotto, E., Spencer, S., Steigerwald, R., Strouse, D., Hansen, S., Filos, A., Brooks, E., Gazeau, M.,
Sahni, H., Singh, S., and Mnih, V . (2022). In-context reinforcement learning with algorithm distillation.
Lewis, P., Perez, E., Piktus, A., Petroni, F., Karpukhin, V ., Goyal, N., Kulkarni, V ., Lewis, M., Reimers, N., Riedel, S., and Zettlemoyer,
L. (2020). Retrieval-augmented generation for knowledge-intensive NLP tasks. In Advances in Neural Information Processing
Systems (NeurIPS).
17

## Page 18

A PREPRINT - J UNE 6, 2025
Nakano, R., Hilton, J., Balaji, S., Wu, J., Ouyang, L., Kim, C., Hesse, C., Jain, S., Kosaraju, V ., Saunders, W., Jiang, X., Cobbe,
K., Eloundou, T., Krueger, G., Button, K., Knight, M., Chess, B., and Schulman, J. (2022). Webgpt: Browser-assisted
question-answering with human feedback.
Naveed, H., Khan, A. U., Qiu, S., Saqib, M., Anwar, S., Usman, M., Akhtar, N., Barnes, N., and Mian, A. (2024). A comprehensive
overview of large language models.
Niven, P. R. and Lamorte, B. (2016). Objectives and Key Results: Driving Focus, Alignment, and Engagement with OKRs. John
Wiley & Sons, Inc.
O’neil, C. (2017). Weapons of math destruction: How big data increases inequality and threatens democracy. Crown.
Peffers, K., Tuunanen, T., Gengler, C. E., Rossi, M., Hui, W., Virtanen, V ., and Bragge, J. (2020). Design science research process:
A model for producing and presenting information systems research.
Peffers, K., Tuunanen, T., Rothenberger, M. A., and Chatterjee, S. (2007). A design science research methodology for information
systems research. Journal of Management Information Systems, 24(3):45–77. Cited as methodological foundation in 2020
contexts.
Pitkäranta, T. and Pitkäranta, L. (2024). Bridging human and ai decision-making with llms: The ragada approach. In Proceedings of
the 26th International Conference on Enterprise Information Systems, Query date, volume 6, page 45.
Shinn, N., Cassano, F., Berman, E., Gopinath, A., Narasimhan, K., and Yao, S. (2023a). Reflexion: Language agents with verbal
reinforcement learning.
Shinn, N., Cassano, F., Berman, E., Gopinath, A., Narasimhan, K., and Yao, S. (2023b). Reflexion: Language agents with verbal
reinforcement learning.
Suleyman, M. and Bhaskar, M. (2023). The Coming Wave: Technology, Power, and the Twenty-First Century’s Greatest Dilemma.
Crown Publishing Group, New York.
Vaswani, A., Shazeer, N., Parmar, N., Uszkoreit, J., Jones, L., Gomez, A. N., Kaiser, Ł., and Polosukhin, I. (2017). Attention is all
you need. Advances in neural information processing systems, 30.
Wang, L., Ma, C., Feng, X., Zhang, Z., Yang, H., Zhang, J., Chen, Z., Tang, J., Chen, X., Lin, Y ., Zhao, W. X., Wei, Z., and Wen, J.
(2024a). A survey on large language model based autonomous agents. Frontiers of Computer Science, 18(6).
Wang, L., Ma, C., Feng, X., Zhang, Z., Yang, H., Zhang, J., Chen, Z., Tang, J., Chen, X., Lin, Y ., Zhao, W. X., Wei, Z., and Wen, J.
(2024b). A survey on large language model based autonomous agents. Frontiers of Computer Science, 18(6).
Weidinger, L., Mellor, J., Rauh, M., Griffin, C., Uesato, J., Huang, P.-S., Cheng, M., Glaese, M., Balle, B., Kasirzadeh, A., Kenton,
Z., Brown, S., Hawkins, W., Stepleton, T., Biles, C., Birhane, A., Haas, J., Rimell, L., Hendricks, L. A., Isaac, W., Legassick,
S., Irving, G., and Gabriel, I. (2021). Ethical and social risks of harm from language models.
Wilson, E. O. (1999). Consilience: The unity of knowledge. Vintage; Reprint edition. ISBN: 978-0679768678.
Wu, Q., Bansal, G., Zhang, J., Wu, Y ., Li, B., Zhu, E., Jiang, L., Zhang, X., Zhang, S., Liu, J., Awadallah, A. H., White, R. W.,
Burger, D., and Wang, C. (2023). Autogen: Enabling next-gen llm applications via multi-agent conversation.
Xi, Z., Chen, W., Guo, X., He, W., Ding, Y ., Hong, B., Zhang, M., Wang, J., Jin, S., Zhou, E., Zheng, R., Fan, X., Wang, X., Xiong,
L., Zhou, Y ., Wang, W., Jiang, C., Zou, Y ., Liu, X., Yin, Z., Dou, S., Weng, R., Cheng, W., Zhang, Q., Qin, W., Zheng, Y ., Qiu,
X., Huang, X., and Gui, T. (2023). The rise and potential of large language model based agents: A survey.
Yampolskiy, R. V . (2014). The AI containment problem.Journal of Consciousness Studies, 23(9–10):222–234. Discusses technical
and organisational safeguards (“containment”) for superintelligent AI systems.
Yang, S., Nachum, O., Du, Y ., Wei, J., Abbeel, P., and Schuurmans, D. (2023). Foundation models for decision making: Problems,
methods, and opportunities.
Yang, Y ., Chai, H., Song, Y ., Qi, S., Wen, M., Li, N., Liao, J., Hu, H., Lin, J., Chang, G., Liu, W., Wen, Y ., Yu, Y ., and Zhang, W.
(2025). A survey of ai agent protocols.
Yao, S., Zhao, J., Yu, D., Du, N., Shafran, I., Narasimhan, K., and Cao, Y . (2023a). React: Synergizing reasoning and acting in
language models.
Yao, S., Zhao, J., Yu, D., Du, N., Shafran, I., Narasimhan, K., and Cao, Y . (2023b). React: Synergizing reasoning and acting in
language models.
18
