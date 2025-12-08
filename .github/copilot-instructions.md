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

## Reviews

- **Content reviews**: Focus on accuracy, clarity, and completeness of technical content
  - Go ahead and make changes for spelling, grammar, style, formatting, and voice
  - For more substantial changes, suggest what should be changed
  - Ensure explanations are correct and easy to understand
  - Verify code examples work as intended
  - Add missing examples

## Terminology

- **Declare vs Define**: Use "declare" when introducing a name/type/signature to
  the compiler, and   "define" when specifying the complete implementation or
  structure. Methods are declared; their bodies are defined. Types are declared;
  their fields/structure define them.

## Commits and pushes

- **Conventional commits**: Use conventional commit messages for all changes
  - Example: "fix: correct slice length calculation"
  - Example: "docs: update chapter on concurrency"
  - Example: "feat: add new chapter on error handling"
