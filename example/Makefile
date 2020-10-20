TF ?= terraform-v0.12.29
TEST_OUT_DIR ?= testout/$(shell cat USE | tr -d '\n' || echo current)
TF_LOG ?= TRACE

.PHONY: dev
dev:
	rm tf.log tf.*.log || true
	rm -rf ./tf.*.pprof || true
	make testdata/01-bootstrap init plan apply
	make testdata/02-shoal plan apply
	make destroy

.PHONY: init
init:
	cd .. ; make build
	mv ../dist/darwin_amd64/terraform-provider-kubectl .terraform/plugins/darwin_amd64/terraform-provider-kubectl
	$(TF) init

.PHONY: testout
testout:
	mkdir -p $(TEST_OUT_DIR)

.PHONY: apply
apply: testout
	TF_LOG=$(TF_LOG) TF_LOG_PATH=$(TEST_OUT_DIR)/apply.log time -l $(TF) apply -auto-approve -var kubeconfig=$(HOME)/.kube/config

.PHONY: destroy
destroy: testout
	TF_LOG=$(TF_LOG) TF_LOG_PATH=$(TEST_OUT_DIR)/destroy.log time -l $(TF) destroy -auto-approve -var kubeconfig=$(HOME)/.kube/config

.PHONY: plan
plan: testout
	TF_LOG=$(TF_LOG) TF_LOG_PATH=$(TEST_OUT_DIR)/plan.log time -l $(TF) plan -var kubeconfig=$(HOME)/.kube/config

STEPS = $(wildcard testdata/*)

.PHONY: $(STEPS)
.SECONDEXPANSION:
$(STEPS):%:$$(wildcard %/*.yaml)
	@echo Preparing $@
	cp $@/*.tf .
	@echo $(shell basename $@) > USE
