#include <string.h>

#include "cJSON.h"
#include "protobuf-c/protobuf-c.h"

#include "esp_err.h"
#include "esp_log.h"

#include "config.h"
#include "context.h"
#include "error.h"
#include "iot.h"
#include "mqtt.h"

static const char *TAG = "iot";

static esp_err_t iot_handle_config(context_t *context, const uint8_t *config, size_t size) {
    ARG_UNUSED(context);
    ARG_UNUSED(config);
    if (size <= 0) {
        return ESP_OK;
    }
    Hydroponics__Config *pbcfg;
    ESP_ERROR_CHECK(config_parse(config, size, &pbcfg));
    ESP_ERROR_CHECK(config_dump(pbcfg));
    ESP_ERROR_CHECK(config_free(pbcfg));
    return ESP_OK;
}

static esp_err_t iot_handle_command(context_t *context, const uint8_t *command, size_t size) {
    ARG_UNUSED(context);
    ARG_UNUSED(command);
    ARG_UNUSED(size);
    // TODO(sobrinho): Add proper handling of commands.
    return ESP_OK;
}

static esp_err_t iot_handle_publish_telemetry(context_t *context, uint8_t **data, size_t *size) {
    ARG_CHECK(data != NULL, ERR_PARAM_NULL);
    ARG_CHECK(size != NULL, ERR_PARAM_NULL);

    *size = 0;
    cJSON *root = cJSON_CreateObject();
    cJSON_AddNumberToObject(root, "sensors.temp.indoor", context->sensors.temp.indoor);
    cJSON_AddNumberToObject(root, "sensors.temp.probe", context->sensors.temp.probe);
    cJSON_AddNumberToObject(root, "sensors.humidity", context->sensors.humidity);
    cJSON_AddNumberToObject(root, "sensors.pressure", context->sensors.pressure);
    cJSON_AddNumberToObject(root, "sensors.ec.value", context->sensors.ec.value);
    cJSON_AddNumberToObject(root, "sensors.ph.value", context->sensors.ph.value);

    char *msg = cJSON_Print(root);
    cJSON_Delete(root);
    if (msg == NULL) {
        ESP_LOGE(TAG, "Failed to produce the json telemetry data");
        return ESP_FAIL;
    }
    *data = (uint8_t *) msg;
    *size = strlen(msg);
    return ESP_OK;
}

static esp_err_t iot_handle_publish_state(context_t *context, uint8_t **data, size_t *size) {
    ARG_CHECK(data != NULL, ERR_PARAM_NULL);
    ARG_CHECK(size != NULL, ERR_PARAM_NULL);

    *size = 0;
    char *msg;
    ESP_ERROR_CHECK(context_get_state_message(context, &msg));
    if (msg != NULL) {
        *data = (uint8_t *) msg;
        *size = strlen(msg);
    }
    return ESP_OK;
}

static mqtt_config_t config = {
        .handle_config = iot_handle_config,
        .handle_command = iot_handle_command,
        .handle_publish_telemetry = iot_handle_publish_telemetry,
        .handle_publish_state = iot_handle_publish_state,
};

esp_err_t iot_init(context_t *context) {
    ESP_ERROR_CHECK(mqtt_init(context, &config));
    return ESP_OK;
}
