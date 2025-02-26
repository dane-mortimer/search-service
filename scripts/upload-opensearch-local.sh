#!/usr/bin/env bash

docs=(
  "Introduction to OpenSearch|OpenSearch is a powerful search engine based on Elasticsearch."
  "Getting Started with Docker|Docker simplifies application deployment using containers."
  "Advanced Go Programming|Learn advanced Go concepts like concurrency and profiling."
  "Data Science with Python|Python is widely used for data analysis and machine learning."
  "Cloud Computing Basics|Understand the fundamentals of cloud computing and AWS."
  "Machine Learning Fundamentals|An introduction to supervised and unsupervised learning."
  "Cybersecurity Essentials|Learn the basics of cybersecurity to protect digital assets."
  "Blockchain Technology|A deep dive into how blockchain enables decentralized systems."
  "DevOps Best Practices|Improve software development workflows using DevOps principles."
  "Kubernetes for Beginners|Container orchestration with Kubernetes explained."
  "RESTful API Design|Best practices for designing scalable RESTful APIs."
  "GraphQL vs REST|Understanding the differences between GraphQL and REST APIs."
  "Microservices Architecture|Breaking down monolithic applications into microservices."
  "Introduction to Rust|Learn why Rust is gaining popularity among systems programmers."
  "SQL vs NoSQL Databases|Key differences between relational and non-relational databases."
  "Artificial Intelligence Trends|Latest developments in AI and machine learning."
  "Python Web Frameworks|Django vs Flask: Choosing the right framework for web development."
  "Big Data Processing|How to handle large-scale data processing using Apache Spark."
  "Cloud Security|Essential practices to secure cloud infrastructure."
  "The Future of Edge Computing|How edge computing is transforming the tech landscape."
  "Software Development Life Cycle|Understanding SDLC methodologies and best practices."
  "Internet of Things (IoT) Basics|Learn how IoT is connecting devices and changing industries."
  "Serverless Computing|Explore the benefits and limitations of serverless architectures."
  "Deep Learning Applications|How deep learning is powering AI innovations."
  "Functional Programming Concepts|Understanding functional programming paradigms."
  "Quantum Computing Overview|A look into the principles of quantum computing."
  "Natural Language Processing|How NLP is transforming human-computer interactions."
  "Automating Tasks with Bash|Writing efficient shell scripts for automation."
  "Agile vs Waterfall|Comparing traditional and modern software development approaches."
  "Hybrid Cloud Strategies|Leveraging hybrid cloud models for optimal performance."
  "Data Warehousing|Understanding data warehousing concepts and best practices."
  "Digital Marketing Trends|Latest trends in digital marketing and SEO strategies."
  "Computer Vision|How machines interpret and process visual data."
  "API Gateway Fundamentals|Managing APIs efficiently with an API gateway."
  "Enterprise Software Architecture|Building scalable and maintainable enterprise applications."
  "Mobile App Development|Key considerations for building mobile applications."
  "Security Best Practices for Developers|How to write secure code and prevent vulnerabilities."
  "Introduction to CI/CD|Understanding Continuous Integration and Continuous Deployment."
  "Data Encryption Techniques|Protecting data using encryption algorithms."
  "Performance Tuning in Databases|Optimizing database performance and query efficiency."
)


for i in "${!docs[@]}"
do
  title="${docs[$i]%%|*}"
  content="${docs[$i]#*|}"
  curl -X POST http://localhost:9200/my-index/_doc/$((i+1)) -H "Content-Type: application/json" -d "{
    \"title\": \"$title\",
    \"content\": \"$content\"
  }"
done


curl -X GET http://localhost:9200/my-index/_search?pretty