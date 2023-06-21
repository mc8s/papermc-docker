FROM amazoncorretto:17-alpine-jdk
LABEL authors="Mc8s"

WORKDIR /app

COPY server.jar /app/
COPY eula.txt /app/

ENTRYPOINT ["java", "-jar", "server.jar"]