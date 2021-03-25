'use strict';

export const DEFAULT_SELECTOR = '.collapsible';

const DEFAULT_OPTIONS = {
    toggleClassName: 'collapsible',
    suffixActiveClass: '--active'
}

class Collapsible {
    /**
     * Initializer
     *
     * @param {Element} elementRef - Link list element containing with each link item inside
     * @param {Object} options - Passing some parameters to override default options
     */
    constructor(elementRef, options) {
        // Initialise attributes
        this.options = Object.assign(DEFAULT_OPTIONS, options);
        this.toggleClassName = `${this.options.toggleClassName}${this.options.suffixActiveClass}`;

        this.elementRef = elementRef;
        this.elementToggle = this.elementRef.querySelector(`.${this.elementRef.dataset.collapsibleToggle}`);

        this._bind();
        this._setup();
    }

    // Event Handlers
    onClickCollapsible() {
        let activeElement = document.querySelector(`.${this.toggleClassName}`);

        this.elementRef.classList.add(this.toggleClassName);

        if (activeElement) {
            activeElement.classList.remove(this.toggleClassName);
        }
    }

    // Private

    /**
     * Bind all functions to the local instance scope.
     * */
    _bind() {
        this.onClickCollapsible = this.onClickCollapsible.bind(this);
    }

    _setup() {
        this.elementToggle.addEventListener('click', this.onClickCollapsible);
    }
}

export default Collapsible;
