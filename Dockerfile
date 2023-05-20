FROM amazoncorretto:17-alpine-jdk
LABEL authors="Jonas-be"

WORKDIR /app

COPY server.jar /app/
COPY eula.txt /app/

ENTRYPOINT ["java", "-jar", "server.jar"]