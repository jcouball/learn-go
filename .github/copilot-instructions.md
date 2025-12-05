## Voice and Style

- **Main body text**: Use second person ("you") with active voice for explanations and instructions
  - Example: "You can modify the slice contents but you can't change its length."

- **Admonitions (notes, tips, warnings)**: Use imperative voice (commands) without "you"
  - Example: "Prefer structs over maps when keys are known at compile time."
  - Example: "Use pointers only when you need to modify the value."

- **Code comments**: Use declarative statements or imperative voice
  - Example: "This changes the original map"
  - Example: "Returns the number of bytes read"

- **Avoid**:
  - Passive voice ("should be used" → "use")
  - Weak phrases ("you might want to consider" → "consider")
  - Starting admonitions with "You should..." (use imperative instead)

## Commits and pushes

- **Conventional commits**: Use conventional commit messages for all changes
  - Example: "fix: correct slice length calculation"
  - Example: "docs: update chapter on concurrency"
  - Example: "feat: add new chapter on error handling"
