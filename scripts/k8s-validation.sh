#!/bin/bash

deploy_chart() {
    echo "--------------Deploying Helm Chart--------------"
    helm upgrade dynamic-pv-scaler ./deploy/helm -f \
    ./deploy/helm/values.yaml --install --namespace dynamic-pv-scaler
}

validate_chart() {
    echo "--------------Testing Helm Chart--------------"
    helm test dynamic-pv-scaler --namespace dynamic-pv-scaler
}

lint_chart() {
    echo "--------------Linting Helm Chart--------------"
    helm lint ./deploy/helm/.
}

validate_container_state() {
    echo "--------------Validating Deployment Status--------------"
    output=$(kubectl get pods -n keycloak -l app=dynamic-pv-scaler \
    -o jsonpath="{.items[*]['status.phase']}")
    if [ "${output}" != "Running" ] && [ "${output}" != "" ]
    then
        echo "Container is not healthy"
        exit 1
    else
        echo "Container is running fine"
    fi
}

main_function() {
    lint_chart
    deploy_chart
    validate_chart
    sleep 30s
    validate_container_state
}

main_function
