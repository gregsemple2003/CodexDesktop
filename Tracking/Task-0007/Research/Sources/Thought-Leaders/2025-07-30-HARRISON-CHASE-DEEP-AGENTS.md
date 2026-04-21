Deep Agents

Products LangSmith Platform Observability See exactly what your agents are doing Evaluation Score and improve agent performance Deployment Ship and scale agents in production Fleet Agents for the whole company Open Source Frameworks deepagents Build long-running agents for complex tasks langchain Quick start agents with any model provider langgraph Build reliable agents with low-level control Learn Resources Blog Customer Stories Guides Max Agency How-To LangChain Academy YouTube Documentation Community LangSmith for Startups Events Community Docs Company About Careers Partners Pricing Try LangSmith Get a demo Try LangSmith Get a demo Go back to blog Harrison's In the Loop Deep Agents

# Deep Agents

Harrison Chase July 30, 2025 X min Create agents Share

Using an LLM to call tools in a loop is the simplest form of an agent. This architecture, however, can yield agents that are “shallow” and fail to plan and act over longer, more complex tasks. Applications like “ Deep Research ”, “ Manus ”, and “ Claude Code ” have gotten around this limitation by implementing a combination of four things: a planning tool, sub agents, access to a file system, and a detailed prompt.

Acknowledgements: this exploration was primarily inspired by Claude Code and reports of people using it for more than just coding . What about Claude Code made it general purpose, and could we abstract out and generalize those characteristics?

## Deep agents in the wild

The dominant agent architecture to emerge is also the simplest: running in a loop, calling tools.

Doing this naively, however, leads to agents that are a bit shallow. “Shallow” here refers to the agents inability to plan over longer time horizons and do more complex tasks.

Research and coding have emerged as two areas where agents have been created that buck this trend. All of the major model providers have an agent for Deep Research and for “async” coding tasks. Many startups and customers are creating versions of these for their specific vertical.

I refer to these agents as “deep agents” - for their ability to dive deep on topics. They are generally capable of planning more complex tasks, and then executing over longer time horizons on those goals.

What makes these agents good at going deep?

The core algorithm is actually the same - it’s an LLM running in a loop calling tools. The difference compared to the naive agent that is easy to build is:

- A detailed system prompt

- Planning tool

- Sub agents

- File system

## Characteristics of deep agents

Detailed system prompt

Claude Code’s recreated system prompts are long. They contain detailed instructions on how to use tools. They contain examples (few shot prompts) on how to behave in certain situations.

Claude Code is not an anomaly - most of the best coding or deep research agents have pretty complex system prompts. Without these system prompts, the agents would not be nearly as deep. Prompting matters still!

Planning tool

Claude Code uses a Todo list tool . Funnily enough - this doesn’t do anything! It’s basically a no-op. It’s just context engineering strategy to keep the agent on track.

Deep agents are better at executing on complex tasks over longer time horizons. Planning (even if done via a no-op tool call) is a big component of that.

Sub agents

Claude Code can spawn sub agents . This allows it to split up tasks. You can also create custom sub agents to have more control. This allows for "context management and prompt shortcuts" .

Deep agents go deeper on topics. This is largely accomplished by spinning up sub agents that specifically focused on individual tasks, and allowing those sub agents to go deep there.

File System

Claude Code (obviously) has access to the file system and can modify files on there, not just to complete its task but also to jot down notes. It also acts as a shared workspace for all agents (and sub agents) to collaborate on.

Manus is another example of a deep agent that makes significant use of a file system for “memory”.

Deep agents run for long periods of time and accumulate a lot of context that they need to manage. Having a file system handy to store (and then later read from) is helpful for doing so.

## Build your deep agent

In order to make it easier for everyone to build a deep agent for their specific vertical, I hacked on an open source package (

deepagents

) over the weekend. You can easily install it with

pip install deepagents

and then read instructions for how to use it here .

This package attempts to create a general purpose deep agent that can be customized to create your own custom version.

It comes with built-in components mapping to the above characteristics:

- A system prompt inspired by Claude Code, but modified to be more general

- A no-op Todo list planning tool (same as Claude Code)

- Ability to spawn sub-agents, and specify your own

- A mocked out “virtual file system” that uses the agents state (a preexisting LangGraph concept)

You can easily create your own deep agent by passing in a custom prompt (will be inserted into the larger system prompt as custom instructions), custom tools, and custom subagents. We put together a simple example of a "deep research" agent built on top of

deepagents

.

TRY OUT

deepagents

HERE

Documentation

### Related content

Agent Architecture Deep Agents Open Source

#### Running Subagents in the Background

Hunter Lovell Colin Francis April 16, 2026 4 min Deep Agents Engineering

#### How We Made Our Docs Test Themselves

Naomi Pentrel April 15, 2026 3 min Harrison's In the Loop

#### Your harness, your memory

Harrison Chase April 11, 2026 7 min Sign up for our newsletter to stay up to date Thank you! Your submission has been received! Oops! Something went wrong while submitting the form.

### See what your agent is really doing

LangSmith, our agent engineering platform, helps developers debug every agent decision, eval changes, and deploy in one click.

Try LangSmith Get a demo

###### Products

LangSmith Platform LangSmith Observability LangSmith Evaluation LangSmith Deployment LangSmith Fleet Deep Agents LangChain LangGraph

###### Resources

Blog Customer Stories Guides LangChain Academy Community Events Changelog Docs Support

###### Company

About Careers Partners Trust Center Marketing Assets

###### Sign up for our newsletter to stay up to date

Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. All systems operational Privacy policy Terms of service
