# CLAUDE.md

Behavioral guidelines to reduce common LLM coding mistakes. Merge with project-specific instructions as needed.

**Tradeoff:** These guidelines bias toward caution over speed. For trivial tasks, use judgment.

## 1. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

Before implementing:

- State your assumptions explicitly. If uncertain, ask.
- If multiple interpretations exist, present them - don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, stop. Name what's confusing. Ask.

## 2. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If you write 200 lines and it could be 50, rewrite it.

Ask yourself: "Would a senior engineer say this is overcomplicated?" If yes, simplify.

## 3. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:

- Don't "improve" adjacent code, comments, or formatting.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, mention it - don't delete it.

When your changes create orphans:

- Remove imports/variables/functions that YOUR changes made unused.
- Don't remove pre-existing dead code unless asked.

The test: Every changed line should trace directly to the user's request.

## 4. Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:

- "Add validation" → "Write tests for invalid inputs, then make them pass"
- "Fix the bug" → "Write a test that reproduces it, then make it pass"
- "Refactor X" → "Ensure tests pass before and after"

For multi-step tasks, state a brief plan:

```
1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]
```

Strong success criteria let you loop independently. Weak criteria ("make it work") require constant clarification.

## 5. Worktree Isolation

**Never edit the main checkout. Always branch from `claude`.**

- Call `EnterWorktree` before the first file edit of any task, naming it `worktree-claude-<topic>`. The name becomes the branch name, so the prefix carries through to the remote.
- Immediately after entering, run `git fetch origin && git reset --hard origin/claude` - the default base ref is `origin/main`, which is not what this project wants.

When the work is done:

1. Commit.
2. Push the branch, then open the pull request against `claude`:

   ```
   git push -u origin HEAD
   tea pr create --base claude --title "<title>"
   ```

   `--base` falls back to the repo default branch, so it is never optional. `--head` defaults to the current branch, so run this from inside the worktree.

3. Report the pull request URL, then stay in the worktree.
4. Act on the next instruction:
   - `merge` - merge and clean up, see below.
   - `abandon` - drop branch remote and local, delete worktree.
   - anything else - the work was not satisfactory. Keep going in the same worktree; new commits pushed to the same branch land in the open pull request automatically.

Never merge unless told to. Never call `ExitWorktree` unless told to.

Always post the PR last for easier browsing and as a reminder for followed up
work in the branch.

Merging and cleaning up, in order:

1. `tea pr merge <index>`
2. `git push origin --delete <branch>` - `tea` leaves the remote branch behind.
3. Delete the worktree and the local branch:
   - Verify the merge landed:
     `git fetch origin && git merge-base --is-ancestor <sha> origin/claude`
   - `ExitWorktree` with `action: "remove"` and `discard_changes: true`. Without
     the flag it refuses — expected, the merge is on the remote, not on local
     `claude`. Never pass the flag before the check above.
   - Confirm nothing survived:
     `git for-each-ref --format='%(refname)' | grep <branch>`
     If it prints anything, `git branch -D <branch>`.
4. Back in the main checkout, `git pull` so `claude` includes the merge.

## 6. TODO file

@TODO.md

Instructions:

- Modifying the TODO file is forbidden, except for ticking existing checkboxes
- When the user paste an element from the TODO file to the agent, tick the
  corresponding checkbox alongside the commit resulting from the work done
- A `##` chapter in the TODO file that has all its checkboxes ticked shall be
  removed, unless its date is the current day.

---

**These guidelines are working if:** fewer unnecessary changes in diffs, fewer rewrites due to overcomplication, and clarifying questions come before implementation rather than after mistakes.
