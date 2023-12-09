FROM amazoncorretto:17-alpine-jdk
LABEL authors="Mc8s"

WORKDIR /app

COPY server.jar /preset/
COPY eula.txt /preset/
COPY start.sh /exec/

ENTRYPOINT ["/exec/start.sh"]
