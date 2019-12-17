#ifndef HYDROPONICS_DRIVER_STORAGE_H
#define HYDROPONICS_DRIVER_STORAGE_H

#include "esp_err.h"

#include "context.h"

#define STORAGE_KEY_WIFI_SSID     "wifi_ssid"
#define STORAGE_KEY_WIFI_PASSWORD "wifi_password"

esp_err_t storage_init(context_t *context);

esp_err_t storage_get_string(const char *key, char **buf, size_t *length);

esp_err_t storage_set_string(const char *key, char *buf);

#endif //HYDROPONICS_DRIVER_STORAGE_H
