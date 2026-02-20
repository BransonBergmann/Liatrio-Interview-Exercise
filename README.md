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

- [ ] Golang  
  - [ ] Fiber  
- [ ] GitHub Actions  
- [ ] Docker  
- [ ] Cloud Platform  
- [ ] GitHub Workflow  

---

## Practical Goals

- [x] GitHub repo created  
- [x] Go & Fiber installed  
- [ ] Simple application started  

---

## Deliverables

Minified JSON file that returns:

- Dynamic timestamp (within a few seconds of request)  
- Text string: `"My name is..."`

---

## Golang & Fiber Resources

- https://medium.com/@andra.gws/building-a-scalable-api-with-go-and-fiber-a-step-by-step-guide-c0fed11db1d2  
  - Go has very high performance & is great at concurrency  
  - Fiber is built on top of Fasthttp  

- https://zetcode.com/golang/fiber/  
  - Covers JSON usage  

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