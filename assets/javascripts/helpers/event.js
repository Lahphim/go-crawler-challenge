'use strict';

/* eslint-disable no-undef */
export const dispatchEventFromElement = (element, eventName, eventDetail = {}) => {
    const event = new CustomEvent(eventName, { detail: eventDetail });
    element.dispatchEvent(event);
};
