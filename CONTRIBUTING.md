# Contributing to Coinspy

Everyone is welcome to contribute to the project. Contributions may be anything that drives the project forward, with

* bug reports,
* suggestions and
* updating code

being the basics.

## :bug: Bug Reports

_TL;DR: Is it really a bug? --> Has nobody else already reported it? --> Does it still exist in the most current version? --> File a helpful issue._

If you encouter something that classifies as a bug for you, please ensure that it really is a bug: A currently implemented function, suggested by the accompanying documentation or the usage message, that is not working as defined.

Should your finding indeed be a bug as by above definition, please have a look at the [open issues][OpenIssues] and check if someone else has already filed an issue regarding the bug you have encountered. If an issue already exists, feel free to comment on it, but do not file a new issue.

If no issue exists that describes the bug you have encountered, please ensure that you are using the most current version of the software that is available. Try the [latest stable version][LatestStable], followed by - if the bug still exists in the latest stable version - the [latest development version][LatestDev] from the master branch.

In case that the latest working version of the software still cotains the bug you have found, please [file an issue][NewIssue]. When creating the new issue, adhere to the following guidelines:

* The title SHOULD already give a hint on what functionality is broken.
* The description MUST contain a description of
  * what you have tried to accomplish and
  * how you wanted to accomplish it.
* That description MUST be extensive and precise enough to reproduce your activities.
* The description MUST contain all relevant version numbers, at least
  * the version of Coinspy.
* The description SHOULD cotain a statement that you have followed above test instructions.
  * For absolute clarity, please include Go version and commit IDs.

After the bug report has been filed, it will be reviewed by the code owner(s). All further communication will be handled via the issue comments.

## :bulb: Suggestions

_TL;DR: Is the functionality missing from the latest development version? --> Has nobody else already suggested it? --> File a helpful issue._

If some specific functionality is missing from the [latest development version][LatestDev] of the software that you would like to see implemented, head over to the issues and review the [issues labeled `suggestion`][IssuesLabeledSuggestion].

In case someone else has already suggested the functionality you are looking for, feel free to comment on it, but do not file a new issue.

Should there be no issue suggesting the functionality you are looking for, go ahead and [file an issue][NewIssue]. When creating the new issue, adhere to the following guidelines:

* The title SHOULD already give a hint on what functionality shall be implemented.
* The description MUST contain a description of what functionality exactly you are looking for.
  * Feel free to go into implementation details like CLI arguments and expected output.
* That description MUST be extensive and precise enough for a regular user to understand your intentions and the outcome.

After the suggestion has been filed, it will be reviewed by the code owner(s). All further communication will be handled via the issue comments.

## :memo: Updating Code

_TL;DR: Fork the repository. --> Develop bugfix/feature in own branch. --> Send merge request against master._

Whether it is a bug or a feature, if you are able to satisfy your own needs by coding you are welcome to directly contribute your code to the project. Start by [forking the repository][ForkRepo]. Once you have your own fork, develop your bugfix/feature in there and finish by sending a merge request.

Here are some general guidelines for contributing code to the project:

* The master branch IS where development starts and ends.
* Your master branch SHOULD always be up-to-date with the upstream master branch.
* Every bugfix/feature SHOULD be developed in its own branch to simplify merge requests.
* Merge requests MUST be filed against the master branch; merge requests that target other branches will be dismissed.

Please keep in mind that every line of code contributed to the project will be licensed under [the project's license][ProjectLicense]. After receiving the merge request, the code owner(s) will review it. All further communication will be handled via the merge request comments.

[OpenIssues]: https://gitlab.com/rbrt-weiler/coinspy/-/issues
[LatestStable]: https://gitlab.com/rbrt-weiler/coinspy/-/tags
[LatestDev]: https://gitlab.com/rbrt-weiler/coinspy/-/tree/master
[NewIssue]: https://gitlab.com/rbrt-weiler/coinspy/-/issues/new
[IssuesLabeledSuggestion]: https://gitlab.com/rbrt-weiler/coinspy/-/issues?label_name%5B%5D=suggestion
[ForkRepo]: https://gitlab.com/rbrt-weiler/coinspy/-/forks/new
[ProjectLicense]: https://gitlab.com/rbrt-weiler/coinspy/-/blob/master/LICENSE.txt
