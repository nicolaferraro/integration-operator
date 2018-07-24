package me.nicolaferraro.kamel.flow;

import com.fasterxml.jackson.dataformat.yaml.YAMLMapper;
import me.nicolaferraro.kamel.flow.api.FlowSpec;
import me.nicolaferraro.kamel.flow.api.IntegrationSpec;
import me.nicolaferraro.kamel.flow.api.StepSpec;
import org.apache.camel.builder.RouteBuilder;
import org.apache.camel.model.RouteDefinition;

import java.io.IOException;
import java.io.Reader;
import java.util.List;

public class YamlRouteBuilder extends RouteBuilder {

    private static YAMLMapper MAPPER = new YAMLMapper();

    private IntegrationSpec integrationSpec;

    public YamlRouteBuilder(Reader reader) {
        try {
            this.integrationSpec = MAPPER.readValue(reader, IntegrationSpec.class);
        } catch (IOException ex) {
            throw new YamlParseException("Cannot parse input data", ex);
        }
    }

    public void configure() {
        if (this.integrationSpec != null) {
            configureFlows(this.integrationSpec.getFlows());
        }
    }

    protected void configureFlows(List<FlowSpec> flows) {
        flows.forEach(this::configureFlow);
    }

    protected void configureFlow(FlowSpec flow) {
        if (flow.getSteps().size() > 0) {

            // Consumer
            StepSpec fromStep = flow.getSteps().get(0);
            RouteDefinition route = from(fromStep.getUri());

            if (flow.getId() != null) {
                route = route.id(flow.getId());
            }

            if (flow.getName() != null) {
                route = route.description(flow.getName());
            }

            for (int i=1; i<flow.getSteps().size(); i++) {
                StepSpec to = flow.getSteps().get(i);
                route = route.to(to.getUri());
            }

        }
    }

}