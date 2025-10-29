# 🚀 Go GitHub Workshop Template

**CHANGE 3 none**

[![CI Go](https://img.shields.io/badge/CI_GO-Github_Actions-orange)](https://github.com/ClemNTTS/workshop-template-go)
![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)(https://go.dev/)
![License](https://img.shields.io/badge/License-MIT-green.svg)

A minimal Go project designed for a hands-on workshop. This template provides a starting point to learn two fundamental DevOps practices: **Branch Protection Rules** and **Automated CI** with GitHub Actions

---

## ✨ Features

- **Minimal Go Structure:** A simple `main` package that imports a function from a `src/utils` sub-package.
- **Unit Tests:** A ready-to-use `calculator_test.go` file to validate the code.
- **CI Ready:** Includes a pre-configured GitHub Actions workflow (`.github/workflows/go-ci.yml`) to run tests automatically on every Pull Request.

---

## 🎯 Workshop Goal: From Zero to Protected CI

This repository is a _template_. The goal is for you (or your students) to clone it, push it to a _new_ repository, and then configure it step-by-step.

### Part 1: Branch Protection Rules

1.  **Clone & Re-push:**

    - Clone this template: `git clone https://github.com/YOUR_USERNAME/workshop-template-go.git`
    - Create a new, empty repository on your own GitHub account (e.g., `my-workshop-demo`).
    - Change the remote to point to your new repo: `git remote set-url origin https://github.com/YOUR_USERNAME/my-workshop-demo.git`
    - Push the code: `git push -u origin main`

2.  **Protect `main`:**

    - In your new repo, go to `Settings` > `Branches`.
    - Add a branch protection rule for `main`.
    - Enable **`Require a pull request before merging`**.
    - (Optional) Enable **`Require approvals (1)`**.

3.  **Test It:** Try to make a change locally on your `main` branch and `git push`. Watch it get rejected! 🚫

4.  **The Correct Way:** Make a new branch (`git checkout -b feature/my-change`), push it, and open a Pull Request.

### Part 2: Automated CI with GitHub Actions

1.  **Observe the Workflow:** The CI file (`.github/workflows/go-ci.yml`) is already in this repo.
2.  **Trigger the Action:** Open a _new_ Pull Request (or push a new commit to your existing one from Part 1).
3.  **Watch it Run:** Go to the "Checks" tab on your PR. You'll see your `CI Go` workflow running the `go test ./...` command. 🤖
4.  **Bonus: Lock the Merge:**
    - Go back to your `main` branch protection rule (`Settings` > `Branches`).
    - Enable **`Require status checks to pass before merging`**.
    - Select the `build-and-test` job from the list.

**Congratulations!** 🎉 Your repository is now protected. You can no longer merge a PR until it has been reviewed _and_ all automated tests have passed.

---

## 💻 Local Development

To run the project and tests locally:

```bash
# Run the main application
go run main.go

# Run the unit tests
go test -v ./...
```
