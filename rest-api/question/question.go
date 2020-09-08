package question

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// ResponseQuestion defines the json of from the /api/v1/question REST API endpoint.
type ResponseQuestion struct {
	Index    int    `json:"index"`
	Question string `json:"question"`
}

// Question is an array of strings that contains all possible questions from the game.
var question = []string{
	"________ is going to solve all our problems.",
	"________ is the only way I can work with my colleagues.",
	"'Everything fails all the time' was referring to ________.",
	"50% of digital transformation projects end in ________.",
	"75% of disciplinary action was because of ________.",
	"Being pinged by PagerDuty because of ________.",
	"Copying code directly from Stack Overflow led to ________.",
	"Dear boss, I need to go to KubeCon because I want to learn about ________.",
	"Every deployment has a sidecar for ________.",
	"For my next trick, I will containerise ________.",
	"I drink to forget ________.",
	"I passed my last tech interview by using mostly ________.",
	"I passed the Certified Kubernetes Administrator exam by mostly studying ________.",
	"I'm no scrum master, but I'm pretty sure your team is suffering from ________.",
	"Installing with `curl xxx | sudo bash` often results in ________.",
	"Introducing a brand new CI/CD platform! It contains ________.",
	"It really is for the best if the operations team don't know about ________.",
	"It's a pity that computer science graduates these days are all getting involved in ________.",
	"Just give ________ to operations, they can worry about the random failures.",
	"Kubernetes the hard way would be so much better if it featured ________.",
	"My favourite book is 'A dummies guide to ________'.",
	"My favourite session at KubeCon was an in-depth discussion on ________.",
	"Our last crash-loop backoff was caused by ________.",
	"Our security strategy is underpinned by ________.",
	"The best story I read on k8s.af involved ________.",
	"The company hackathon was completely ruined by ________.",
	"The Gartner magic quadrant for ________.",
	"The highlight of this weeks technology demo was ________.",
	"The real reason we don't release on Fridays.",
	"The team are really struggling with ________.",
	"They said we were crazy. They said we couldn't containerise  ________. They were wrong!",
	"This isn't a witch hunt, we just want to know who was stupid enough to implement ________.",
	"Troubleshooting a Kubernetes deployment using ________.",
	"We complete all our documentation using ________.",
	"We created dashboards in ________ so now we can visualise all our problems.",
	"We have a ________ first strategy.",
	"We implemented ________ because it looks great on my CV.",
	"We implemented a service mesh because of ________.",
	"We remediate vulnerabilities by implementing ________.",
	"We secretly filmed you working on ________, pay us 5 bitcoin or we'll release the recording.",
	"We've added ________ into DevOps.",
	"What idea did you bring back from KubeCon?",
	"What is a developer's best friend?",
	"What makes you uncontrollably swear at the computer?",
	"What makes your job exciting?",
	"What never fails to liven up a retrospective?",
	"What would your CISO find disturbing, yet oddly satisfying?",
	"What's currently sending me 1000 alerts per second?",
	"What's your secret power?",
	"When Google first created Borg, they actually intended it for ________.",
	"Someone left an endpoint open to the internet, now we have lots of ________.",
	"We have a ________ strategy, it's going to solve all our problems",
	"I got 99 problems, but ________ ain't one",
	"All I see is ________ everywhere",
	"Coming to the CNCF this year, ________",
	"Our containers are riddled with bugs so we're going to implement ________.",
	"We migrated all our applications to ________ to solve our problems!",
	"The NSA are now deploying ________ in order to secure everything.",
	"We finally managed to refactor our old applications into ________.",
	"I don't solve problems with ________, I just kill the container.",
	"We implemented ________ and now we're posting a long story on k8s.af",
	"Our main tool for passing an impromptu security audit is ________.",
	"Imagine if the Delorean in Back to the Future had been powered by ________.",
	"The banking industry was forever changed when we first introduced them to ________.",
	"Kubernetes is a slippery slope that leads to ________.",
	"What is infrastructures worst enemy?",
	"Instead of coal, Father Christmas now gives bad DevOps teams ________.",
	"The security team are currently crying because our project deployed ________?",
	"The DevOps team are crying because the CTO decided we will be exclusively using ________?",
	"I'm sorry, I couldn't complete that piece of work because of ________.",
	"My next hit novel will include the exciting drama of how a company implemented ________.",
	"The security team is hiding from me because I keep using ________.",
	"I never truly understood DevSecOps until I encountered ________.",
	"Why is your CPU fan constantly at 100%?",
	"I won't produce any work this week as I'll be focused on ________.",
	"As a DevOps team lead, I'm no stranger to ________.",
	"A canary release is basically just doing ________.",
	"We're using a ________ framework.",
	"When we were at the design stage, the fundamental question we asked was ________.",
	"This application is failing because of ________.",
	"My favourite part of working from home is  ________.",
	"What are you scared to tell your head of security about?",
	"We have a cultural problem, we have too many / too much ________.",
	"We implemented two-factor authentication by implementing  ________.",
	"After troubleshooting an intermittent problem for weeks, we finally discovered the root cause was ________.",
	"To support the applications when they move to production, we use ________.",
	"We don't like image scanning our containers because of ________.",
	"Instead of doing the long overdue documentation, you're currently working on ________.",
	"If you had a university degree in ________ you'd be more hireable.",
	"What is LD_PRELOAD really good for?",
	"When everything is on fire, I grab ________ before running out of the building.",
	"To improve our security, we decided to store all our sensitive data in ________.",
	"The question of 'Will it scale?' was finally answered with ________.",
	"________ over my dead body.",
	"I had to put ________ in my container to make it working.",
	"You missed something if you haven't used the ________ operator.",
	"I'm sure ________ will reimplement ________ soon.",
	"We deliver ________-class products.",
	"Instead of a firewall, we implemented ________.",
	"We store all the SSH keys in /www/______.",
	"Can we go-live with ______ by tomorrow?",
	"Oh yeah, we pushed ______ to production yesterday! Didn't you know?",
	"My Inbox filters ______ emails to Trash.",
}

// GetRandomQuestionV1 is the entry point for GET /api/v1/question
func GetRandomQuestionV1(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	questionIndex := rand.Intn(len(question))
	randomQuestion := question[questionIndex]
	response := ResponseQuestion{Index: questionIndex, Question: randomQuestion}
	log.Printf("Index = %d. Random question: %s", response.Index, response.Question)
	json.NewEncoder(w).Encode(response)
}
