FROM eclipse-temurin:17

RUN apt update && apt upgrade -y && apt install -y git

WORKDIR /

# Clone petclinic repo
RUN git clone https://github.com/spring-projects/spring-petclinic.git

WORKDIR /spring-petclinic

# Build the app
RUN ./gradlew bootJar

# Download the otel java agent
RUN wget https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar

# Download rule_based_routing sampler
RUN wget https://repo1.maven.org/maven2/io/opentelemetry/contrib/opentelemetry-samplers/1.40.0-alpha/opentelemetry-samplers-1.40.0-alpha.jar

CMD java -javaagent:opentelemetry-javaagent.jar -Dotel.javaagent.extensions=./opentelemetry-samplers-1.40.0-alpha.jar -jar build/libs/*.jar --logging.level.org.hibernate.SQL=DEBUG
