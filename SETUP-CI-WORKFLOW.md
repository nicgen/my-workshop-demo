## Part 2: Automating Tests with GitHub Actions 🤖

Now that we are forcing Pull Requests (PRs), we need to make sure those PRs don't break our code _before_ we merge them.

This is called **Continuous Integration (CI)**. The tool for this inside GitHub is called **GitHub Actions**.

### Step 2.1: Understanding the "Workflow"

A "Workflow" is just a simple text file written in `YAML` that describes a series of actions to run.

Our Goal: **Every time a PR is opened that targets `main`, automatically run our `go test ./src/tests/` command to verify that everything works.**

### Step 2.2: Create the Workflow File

GitHub Actions looks for its instructions in a very specific folder at the root of your project:

1.  At the root of your project (locally), create a folder named `.github`.
2.  Inside that folder, create another folder named `workflows`.
3.  Finally, inside `.github/workflows/`, create a new file. Let's call it `go-ci.yml`.

Your structure should look like this:

```

MY-PROJECT/
├── .github/
│   └── workflows/
│       └── go-ci.yml   \<-- This is it\!
├── src/
│   ├── tests/
│   └── utils/
├── go.mod
└── main.go

```

### Step 2.3: Building the Workflow (the `go-ci.yml` file)

Open your empty `go-ci.yml` file. We'll build it piece by piece.

**1. Give our Action a Name**

This is just to looks nice in the GitHub UI.

```yaml
# Name of the workflow (visible in the "Actions" tab)
name: CI Go
```

**2. Define the Trigger (When?)**

We want this action to run _only_ on a Pull Request that targets the `main` branch. In YAML, that's written like this:

```yaml
# ... (the name you put above) ...

# "on:" defines the trigger
on:
  # The event is the "pull_request"
  pull_request:
    # We only target PRs aimed at the "main" branch
    branches: [main]
```

**3. Define the Tasks (What?)**

Now, we define what to do. We create a `job`. Let's give it a name (e.g., `build-and-test`). This `job` will run on a virtual machine (a "runner"). Let's use a standard Linux machine (`ubuntu-latest`).

```yaml
# ... (what we wrote above) ...

# "jobs:" lists all the tasks to do
jobs:
  # The name of our job (you can call it whatever you want)
  build-and-test:
    # Virtual machine this job will run on
    runs-on: ubuntu-latest
```

**4. Define the Steps (How?)**

Our `build-and-test` job needs several `steps` to work:

1.  **Check out the code** from our PR.
2.  **Install the Go environment** on the virtual machine.
3.  **Run the tests**.

Let's translate this to YAML:

```yaml
# ... (what we wrote above) ...

# "steps:" is the sequence of actions for this job
steps:
  # Step 1: Check out the PR's code
  # We use a pre-made "action" from GitHub
  - name: Checkout code
    uses: actions/checkout@v4

  # Step 2: Set up the Go environment
  # We use another public action to install Go
  - name: Set up Go
    uses: actions/setup-go@v5
    with:
      go-version: "1.22" # (or whatever version you are using)
```

**5. The Final Step: Run the tests\!**

This one is for you. How would you run the tests from your command line terminal?
...

Exactly\! The command is `go test -v ./...`.

Let's add this final step to our file. It's a simple `run` command:

```yaml
# ... (all the previous steps) ...

# Step 3: Run tests
# This is our own command
- name: Run tests
  run: go test -v ./src/tests/
```

> Your **go-ci.yml** file is complete\!

### Step 2.4: Test the Workflow

1.  Create a new branch: `git checkout -b add-ci`
2.  Commit your file:
    - `git add .github/workflows/go-ci.yml`
    - `git commit -m "feat: Add CI workflow"`
3.  Push your branch: `git push --set-upstream origin add-ci`
4.  Go to GitHub and **open a Pull Request** from `add-ci` to `main`.

**Watch the magic happen\!** 🤩

On the PR page, you should see a new "Checks" section. Your "CI Go" action is running. If everything goes well, you'll see a nice green checkmark ✅.

If not, there is some errors, take a look at Github Actions logs and fix your code !

---

## 🏆 Bonus: Locking the Merge

Right now, our test runs, but we can still merge the PR even if the tests fail (red ❌). We're going to block that.

1.  Go back to `Settings` \> `Branches`.
2.  Click `Edit` on your `main` rule.
3.  Check the box that says **`Require status checks to pass before merging`**.
4.  A list will appear. You should see the name of your job in it: `build-and-test`. Check that box\!
5.  Click `Save changes`.

### The Final Victory

From now on, the "Merge" button on a Pull Request will be **disabled** until:

1.  It has received the required number of approvals (Part 1).
2.  The GitHub Action (`build-and-test`) has finished successfully (Part 2).

You have now set up a protected and automated development cycle. Well done\!
