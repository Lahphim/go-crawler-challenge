'use strict';

import  LinkCounter, { DEFAULT_SELECTOR as LINK_COUNTER_SELECTOR } from "Components/LinkCounter";

const SELECTORS = {
    screen: 'body.report.show'
};

class ReportScreen {
    constructor() {
        this.linkCounter = document.querySelector(LINK_COUNTER_SELECTOR)

        this._setup();
    }

    // Private

    _setup() {
        this._setupLinkCounter();
    }

    _setupLinkCounter() {
        if (this.linkCounter !== null) {
            new LinkCounter(this.linkCounter);
        }
    }
}

let isReportScreen = document.querySelector(SELECTORS.screen) !== null;

if (isReportScreen) {
    new ReportScreen();
}

