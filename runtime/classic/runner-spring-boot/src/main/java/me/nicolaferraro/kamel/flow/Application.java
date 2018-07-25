package me.nicolaferraro.kamel.flow;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;

@SpringBootApplication
public class Application {

    private static final Logger LOG = LoggerFactory.getLogger(Application.class);

    private static final String INTEGRATION_FILE_NAME = "/etc/camel/integration";

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Configuration
    public static class RoutesConfig {

        @Bean
        public YamlRouteBuilder yamlRouteBuilder() throws IOException {


            LOG.info("Loading integration from: {}", INTEGRATION_FILE_NAME);

            YamlRouteBuilder builder;
            try (Reader reader = new FileReader(INTEGRATION_FILE_NAME)) {
                builder = new YamlRouteBuilder(reader);
            }

            return builder;
        }

    }

}
