#ifndef HYDROPONICS_SENSORS_TEMPERATURE_H
#define HYDROPONICS_SENSORS_TEMPERATURE_H

#include "esp_err.h"

extern float temperature;

esp_err_t temperature_init(void);

#endif //HYDROPONICS_SENSORS_TEMPERATURE_H
