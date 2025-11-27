# pltrends
Investigating trends in programming languages using GitHub commits.

## Data

### Types and top level structure
There are 15 event types reflected in the data, they are:
-pushEvent
-createEvent
-deleteEvent
-issueCommentEvent
-pullRequestEvent
-forkEvent
-watchEvent
-issuesEvent
-gollumEvent
-pullRequestReviewEvent
-commitCommentEvent
-releaseEvent
-memberEvent
-publicEvent
-pullRequestReviewCommentEvent

All of these share the same top level map keys:
["id", "type", "actor", "repo", "payload", "public", "created_at"]

create, issueComment, pullRequest, gollum, pullRequestReview, and pullRequestReviewComment events have "org" as an additional top level map key.

For the purpose of this repo, only the pullRequestEvent type is of interest. Relevant data in pullRequestEvent includes:
- repo.id
- language
- stargazers_count
- watchers_count

### Data retrival
[GH Archive](https://gharchive.org) packages event data into 1 hour increments. For each hour of data, the size of the uncompressed file is around 400MB - 700MB. Extracting only the repo.id, language, stargazers_count, and watchers_count attributes results in a ~99.98% space reduction.

Ultimately, the file sizes will be smaller because there will often be multiple events for the same repo, and only the most recent entry per month will be kept.

pullRequestEvent.
    repo.id
    payload.pull_request.head.repo.
        language
        stargazers_count
        watchers_count
