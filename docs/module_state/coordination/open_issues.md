# Coordination Open Issues

Layer: `3`

Thread: `Coordination`

Status: compatibility mirror of `docs/module_state/coordination.md`.

- CLI implementation is deferred until after 1.0.
- Codex Skill / Plugin packaging remains future work.
- Standard package dry-run has only been performed against a local sample, not a
  real external user project.
- Self-evolution dry-run has produced recommendations but those recommendations
  have not been applied.
- Dogfood / Validation remains useful for recording external validation
  evidence, but it no longer owns the self-evolution loop.
- Self-evolution round 1 has produced
  `docs/anchor_pm/reports/self_evolution_round_1.md`; its recommendations have
  not been applied.
- Thread initialization prompts now live in
  `docs/anchor_pm/thread_initialization.md` so users do not need to invent
  thread setup prompts.
- Final user-facing thread prompts must be complete; users should not have to
  fill in placeholders that Codex can generate.
- Target-project integration now has a copy-paste `Thread Management` prompt at
  `docs/anchor_pm/thread_management_install_prompt.md` with the development
  package path.
- Thread Management installer replies should follow the user's usual language;
  project docs may remain English unless requested otherwise.
- Product Manager thread now owns user operation flows and experience
  optimization.
- Reanchor Detector Core now owns the Contract Version / Reanchor State
  Detector as a core code subsystem.
