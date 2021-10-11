SWGR_DIR := $(PROJECT_ROOT)/api/rest/swagger
SWGR_SPEC_NAME := swagger.yml
SWGR_SPEC_FILE := $(SWGR_DIR)/$(SWGR_SPEC_NAME)

SWGR_SRV_PORT := 18080

SWGR_MK_PREFIX := "Swagger:"

## swagger.generate: Generate swagger specification
.PHONY: swagger.generate
swagger.generate: tools.verify.swagger
	@echo "=======> $(SWGR_MK_PREFIX) generating swagger specification"
	@swagger generate spec -o $(SWGR_SPEC_FILE)

## swagger.serve: Serve the swagger specification as a webpage
.PHONY: swagger.serve
swagger.serve: tools.verify.swagger
	@echo "=======> $(SWGR_MK_PREFIX) serving swagger specification at http://localhost:$(SWGR_SRV_PORT)/docs"
	@swagger serve -F=redoc --port $(SWGR_SRV_PORT) $(SWGR_SPEC_FILE)
