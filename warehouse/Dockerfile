FROM openjdk:12-alpine

ARG JAR_FILE=target/*.jar
COPY ${JAR_FILE} app.jar

CMD ["java","-jar","/app.jar"]
