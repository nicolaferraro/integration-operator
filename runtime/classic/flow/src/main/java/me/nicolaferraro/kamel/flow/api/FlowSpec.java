package me.nicolaferraro.kamel.flow.api;

import java.util.LinkedList;
import java.util.List;

public class FlowSpec {

    private String id;

    private String name;

    private List<StepSpec> steps = new LinkedList<>();

    public FlowSpec() {
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public List<StepSpec> getSteps() {
        return steps;
    }

    public void setSteps(List<StepSpec> steps) {
        this.steps = steps;
    }

}
