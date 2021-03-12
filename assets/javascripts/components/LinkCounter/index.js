'use strict';

export const DEFAULT_SELECTOR = 'div.link-counter';

const DEFAULT_OPTIONS = {
    toggleActiveAt: '.link-counter__title',
    activeClass: 'link-counter__item--active'
}

class LinkCounter {
    /**
     * Initializer
     *
     * @param {Element} elementRef - Link list element containing with each link item inside
     */
    constructor(elementRef) {
        // Initialise attributes
        this.options = DEFAULT_OPTIONS;

        // Root alert container
        this.elementRef = elementRef;
        this.toggleActiveElementList = this.elementRef.querySelectorAll(this.options.toggleActiveAt);

        this._bind();
        this._setup();
    }

    // Event Handlers

    /**
     * Click on a title link to toggle a class named `link-counter__item--active`
     */
    onClickLinkList(e) {
        let parentElement = e.target.parentElement;
        let activeElement = this.elementRef.querySelector(`.${this.options.activeClass}`);

        if (!parentElement.classList.contains(this.options.activeClass)) {
            parentElement.classList.add(this.options.activeClass);
        }

        if (activeElement) {
            activeElement.classList.remove(this.options.activeClass);
        }
    }

    // Private

    /**
     * Bind all functions to the local instance scope.
     * */
    _bind() {
        this.onClickLinkList = this.onClickLinkList.bind(this);
    }

    _setup() {
        this.toggleActiveElementList.forEach(element => {
           element.addEventListener('click', this.onClickLinkList);
        });
    }
}

export default LinkCounter;
