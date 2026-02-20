# Liatrio-Interview-Exercise

Start of the project 2-20-26 {
    Legend:
        n = not at all/very little
        s = somewhat/started research
        g = good amount completed
        c = completed/adequate understanding for application

    Research Goals:
        [s] Golang
            [s] Fiber
        [n] Github Actions
        [n] Docker
        [n] Cloud platform
        [n] Github workflow

    Practical Goals:
        [y] Github repo
        [y] Go & Fiber installed
        [n] Simple application started
        [n] 

    Deliverables:
        Minified JSON File that returns:
            -Dynamic timestamp that returns within a few seconds of request
            -Text string of: "My name is..."
}

Golang & Fiber {
    https://medium.com/@andra.gws/building-a-scalable-api-with-go-and-fiber-a-step-by-step-guide-c0fed11db1d2
        -Go has very high performance & great at concurrency
        -Fiber is built on top of Fasthttp
    https://zetcode.com/golang/fiber/
        -talks about JSON specifically
}

Github Actions {
    https://docs.github.com/en/actions/get-started/understand-github-actions
        -CI/CD platform to build & test every pull request to the repo or deploy merged pull requests to prod
        -Workflows
            -Configurable automated process that runs jobs
            -Defined by YAML file in /workflows directory in repo
        -Events
            -Specific activity in a repo that triggers a workflow to run.
        -Jobs
            -Set of steps in a workflow that is executed in the same runner
            -Occur in order & are dependent on each other
        -Actions
            -Pre-Defined & reusable set of jobs/code that perform specific tasks
            -**NOTE**: https://github.com/liatrio/github-actions/tree/master/apprentice-action
                -**INCLUDE THIS IN THE WORKFLOW**
        -Runners
            -Server that runs workflows when triggered.
            -One job at a time.
}