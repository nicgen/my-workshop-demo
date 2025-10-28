### Configure the Branch Protection Rule 🛡️

We will configure GitHub to protect our `main` branch. This will make Pull Requests the _only_ way to get code into `main`.

---

1.  On your GitHub repository page, click the **`Settings`** tab.

2.  In the settings menu on the left, click on **`Branches`**.

3.  You should see a section named "Branch protection rules". Click the **`Add rule`** button.

4.  Name the **Ruleset Name** as you want

5.  Switch **Enforcement status** on **Active**

>     You can ignore Bypass list, it's to setup people who can bypass this rule.

6. Select **Add target** on the _Target branches_ module

- Chose 'include by patterns'
- Enter your default branch (here we use `main`)

---

<br>

7. This is the most important step: Check the box that says **Require a pull request before merging**.

   > This is the setting that will block our direct `git push` attempts.

<br>

8. Let the `Block forces pushes` box checked

9. **Optional** : Also check the box **`Require approvals`**. By default, it will be set to `1`. This means a PR can't be merged (even by its author) until _at least one other person_ reviews and approves it.

10. Scroll to the bottom and click the green **`Create`** button to save the rule.

---

Your `main` branch is now protected! Now, let's try to break it.
