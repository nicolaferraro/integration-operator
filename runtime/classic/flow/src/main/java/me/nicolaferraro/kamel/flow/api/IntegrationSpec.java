package me.nicolaferraro.kamel.flow.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.util.LinkedList;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class IntegrationSpec {

    private int replicas;

    private List<FlowSpec> flows = new LinkedList<>();

    public IntegrationSpec() {
    }

    public int getReplicas() {
        return replicas;
    }

    public void setReplicas(int replicas) {
        this.replicas = replicas;
    }

    public List<FlowSpec> getFlows() {
        return flows;
    }

    public void setFlows(List<FlowSpec> flows) {
        this.flows = flows;
    }

}
