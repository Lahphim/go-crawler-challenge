'use strict';

import Collapsible, { DEFAULT_SELECTOR as COLLAPSIBLE_SELECTOR } from 'Components/Collapsible';

const SELECTORS = {
    screen: 'body.report.show'
};

class ReportScreen {
    constructor() {
        this.collapsibleList = document.querySelectorAll(COLLAPSIBLE_SELECTOR)

        this._setup();
    }

    // Private

    _setup() {
        this._setupCollapsible();
    }

    _setupCollapsible() {
        this.collapsibleList.forEach(collapsible => {
            new Collapsible(collapsible, { activeClass: 'link-counter__item' });
        })
    }
}

let isReportScreen = document.querySelector(SELECTORS.screen) !== null;

if (isReportScreen) {
    new ReportScreen();
}
