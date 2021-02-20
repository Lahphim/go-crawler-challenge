'use strict';

import Alert, { DEFAULT_SELECTOR as ALERT_SELECTOR } from "Components/alert";

document.querySelectorAll(ALERT_SELECTOR).forEach(alert => {
    new Alert(alert);
})
