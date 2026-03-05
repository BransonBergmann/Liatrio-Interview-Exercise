# Liatrio Interview Exercise

**Start Date:** 2-20-26  

---

## Legend

- `n` = not at all / very little  
- `s` = somewhat / started research  
- `g` = good amount completed  
- `c` = completed / adequate understanding for application  

---

## Research Goals

- `c` Golang  
  - `c` Fiber  
- `c` GitHub Actions  
- `c` Docker 
- `c` Cloud Platform  

---

## Practical Goals

- [x] GitHub repo created  
- [x] Go & Fiber installed  
- [x] Simple application started  
- [x] Dockerfile written
- [x] Cloud Service selected
- [x] Github Actions setup
- [x] ACR Storage Set up
- [x] Azure Entra OIDC
- [x] Deployed to Azure Container Apps
---

## Deliverables

Minified JSON file that returns:

- Dynamic timestamp (within a few seconds of request)  
- Text string: `"My name is..."`
- Image tag, git commit, & container revision

---

## Golang & Fiber Resources

- https://medium.com/@andra.gws/building-a-scalable-api-with-go-and-fiber-a-step-by-step-guide-c0fed11db1d2  
  - Go has very high performance & is great at concurrency  
  - Fiber is built on top of Fasthttp  

- https://zetcode.com/golang/fiber/  
  - Covers JSON usage  

Useful Commands:
  - lsof -i :<port#>
    - checks to see if the given port number is being used, and if it is, it gives the Pid of the process
  - kill -9 <Pid>
    - stops the process with the given Pid, useful for opening up local tests
  - go run .
    - runs the local .go file

---

## GitHub Actions

Official Docs:  
https://docs.github.com/en/actions/get-started/understand-github-actions  

### Key Concepts

- **CI/CD Platform**  
  - Builds & tests every pull request  
  - Can deploy merged pull requests to production  

- **Workflows**  
  - Configurable automated processes  
  - Defined by YAML file in `.github/workflows/`  

- **Events**  
  - Activities that trigger workflows  

- **Jobs**  
  - Set of steps executed in the same runner  
  - Run in order and can depend on each other  

- **Actions**  
  - Pre-defined & reusable job components  

  > Include this in your workflow:  
  > https://github.com/liatrio/github-actions/tree/master/apprentice-action  

- **Runners**  
  - Server that runs workflows  
  - One job at a time  

---

## Docker

Useful Commands:
  - docker stop $(docker ps -q)
    - Stops all running docker images
  - docker build --tag <name> .
    - NOTE: Don't forget the dot at the end.
    - builds an image w/ the name you wish
  - docker image ls
    - shows you all avaliable images
  - docker run <name>
    - runs the docker image w/ the name provided
  - docker run --rm -p 8080:8080 echo
    - use this run as it actually sets up the server at port 8080


---

## Azure

Useful Commands:
  - `curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash`
    - actually installs the Azure CLI tools needed to run az commands locally
  - `az acr login --name <acr-name>`
    - authenticates your local Docker client against ACR
  - `az containerapp update --name <app> --resource-group <rg> --image <image>`
    - updates a running Container App to use a new image

---

### OIDC Authentication

- No passwords or secrets stored — uses temporary trust-based tokens
- Requires an App Registration in Entra ID with a Federated Identity Credential pointed at the GitHub repo
- Three values needed: `AZURE_CLIENT_ID` (permissions), `AZURE_TENANT_ID` (directory location), `AZURE_SUBSCRIPTION_ID` (which account)

---

### GitHub Actions Workflow

- 
  - **build-and-push job**: logs into Azure via OIDC, builds the Docker image, pushes to ACR, scans with Trivy, writes a summary
  - **deploy job**: runs only after build-and-push succeeds, updates the Container App with the new image
  - Images are tagged with both the run number (immutable) and `latest`
  - Trivy scan is observe-only — reports CRITICAL/HIGH CVEs to the GitHub Security tab without blocking the pipeline
  - Build summary viewable under the Actions tab → specific run → Summary page