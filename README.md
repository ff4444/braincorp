# Brain Corp
Examples for Brain Corp demonstrating:

- Git and GitHub experience
- Jenkins for CI/CD to go from git to live cloud deployment
  - Google Cloud Build could be used instead
- API design
  - Could manage this in Apigee API Management
  - Swagger is a useful tool as well
- Security and encryption
- Google Cloud Anthos for Jenkins and build/test containers
- Google Kubernetes Engine to configure containers
- Google Cloud Functions in Python (Robot event reporting API)
- Google App Engine design (Go batch processing of queue of events using Dataproc API)
  - Could use Spark also for data processing
  - Could use Google Cloud AI for data inference and summarization
- Google Big Query
  - Could use Google Cloud Storage and Snowflake instead

```mermaid
sequenceDiagram
    participant GitHub
    participant GKE
    participant Jenkins
    participant BuildTestContainer
    participant Google Cloud Functions
    participant Google App Engine
    participant Google Dataproc
    participant BigQuery
    GitHub->>GKE: Container definitions
    GKE->>Jenkins: Instantiate
    GKE->>BuildTestContainer: Instantiate
    Jenkins->>GitHub: Poll
    Jenkins->>BuildTestContainer: Build Python and Go code
    BuildTestContainer->>Jenkins: Test results
    Jenkins->>Google Cloud Functions: Deploy Python API endpoints
    Jenkins->>Google App Engine: Deploy Go robot batch processing app
    Google App Engine->>Google Dataproc: Clean / process
    Google Dataproc->>BigQuery: Robot data storage
```
