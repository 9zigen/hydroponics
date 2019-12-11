#include <string.h>

#include "driver/i2c.h"

#include "esp_log.h"
#include "esp_err.h"

#include "context.h"
#include "error.h"
#include "driver/ezo.h"

#define EZO_EC_ADDR 0x64                /*!< Slave address for Atlas EZO EC module. */
#define SAMPLE_PERIOD 1000              /*!< Reading sample period. */

static const char *TAG = "ezo_ec";
static ezo_sensor_t ec = {
        .type = "EC",
        .probe = "CS150",
        .address = EZO_EC_ADDR,
        .cmd_device_info = {.cmd = "I", .delay_ms = 300, .has_read = true},
        .cmd_read = {.cmd = "R", .delay_ms = 600, .has_read = true},
        .cmd_read_temperature = {.cmd = "RT", .delay_ms = 600, .has_read = true},
        .cmd_status = {.cmd = "S", .delay_ms = 600, .has_read = true},
};

static void ezo_ec_task(void *arg) {
    context_t *context = (context_t *) arg;
    ARG_ERROR_CHECK(context != NULL, ERR_PARAM_NULL);
    ESP_ERROR_CHECK(ezo_init(&ec));

    float last_temp_water = 25.0f;
    float value;
    while (1) {
        TickType_t last_wake_time = xTaskGetTickCount();

        float temp_water = context->sensors.temp.water;
        if (last_temp_water != temp_water && CONTEXT_VALUE_IS_VALID(temp_water)) {
            last_temp_water = temp_water;
            ESP_ERROR_CHECK(ezo_read_temperature_command(&ec, &value, last_temp_water));
        } else {
            ESP_ERROR_CHECK(ezo_read_command(&ec, &value));
        }
        ESP_LOGI(TAG, "EC %.2f uS/cm", context->sensors.ec.value);
        ESP_ERROR_CHECK(context_set_ec(context, value));

        vTaskDelayUntil(&last_wake_time, pdMS_TO_TICKS(SAMPLE_PERIOD));
    }
}

esp_err_t ezo_ec_init(context_t *context) {
    xTaskCreatePinnedToCore(ezo_ec_task, "ezo_ec", 2048, context, 10, NULL, tskNO_AFFINITY);
    return ESP_OK;
}
