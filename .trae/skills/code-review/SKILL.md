---
name: "code-review"
description: "Performs structured code review for bugs, regressions, and test gaps. Invoke when user asks for review, PR feedback, or pre-merge risk assessment."
---

# Code Review

## Purpose

Perform a practical, risk-focused code review. Prioritize correctness, behavior changes, regressions, security, and missing tests over style nitpicks.

## When To Invoke

Invoke this skill when:

- User asks for a code review, PR review, or quality assessment.
- User requests a "review" without detailed implementation work.
- Changes are ready for merge and need risk validation.

Do not invoke this skill for pure feature implementation unless the user also asks for a review.

## Review Workflow

1. Understand scope:
   - Identify changed files and their functional purpose.
   - Clarify assumptions if requirements are ambiguous.
2. Inspect high-risk areas first:
   - Business logic and state transitions.
   - Error handling and edge cases.
   - Security-sensitive paths (auth, permissions, secrets, input validation).
   - Concurrency, transactions, and data consistency.
3. Evaluate behavior impact:
   - Detect breaking changes and compatibility risks.
   - Check performance hotspots introduced by the changes.
4. Assess tests:
   - Verify coverage for new/changed behavior.
   - Flag missing tests for critical branches and failure paths.
5. Report findings:
   - List findings first, ordered by severity.
   - Include concrete evidence (file + symbol or snippet reference).
   - Keep summaries brief and secondary.

## Severity Levels

- Critical: likely production outage, data loss/corruption, major security issue.
- High: significant functional bug or regression risk.
- Medium: correctness or maintainability issue with moderate impact.
- Low: minor issue with low risk.

## Output Format

Use this structure:

1. Findings (ordered by severity):
   - `[Severity]` short title
   - Why it matters
   - Evidence (file/symbol reference)
   - Suggested fix
2. Open Questions / Assumptions (if any)
3. Brief Summary
4. Residual Risks / Test Gaps

If no issues are found, explicitly state:

- "No critical findings."
- Remaining risks or confidence limits (for example, missing integration tests or unverified runtime paths).

## Review Heuristics

- Prefer actionable issues over generic advice.
- Avoid style-only comments unless style causes defects.
- Focus on user-visible behavior and failure modes.
- Ensure recommendations are minimal and practical.
