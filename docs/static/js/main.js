/**
 * CAP TOC - Main JavaScript
 * Contains utility functions and common features for the learning website
 */

document.addEventListener("DOMContentLoaded", function () {
    // Initialize any common UI elements
    initDropdowns();

    // Add formatYear function for templates
    if (typeof window.formatYear !== "function") {
        window.formatYear = function (dateStr) {
            return new Date(dateStr).getFullYear();
        };
    }

    // Parse options from string (for nguphap template)
    if (typeof window.parseOptions !== "function") {
        window.parseOptions = function (optionsStr) {
            try {
                // Try to parse JSON string
                return JSON.parse(
                    optionsStr.replace(/'/g, '"')
                );
            } catch (e) {
                // Fallback to simple comma separation if JSON parse fails
                return optionsStr
                    .split(",")
                    .map((opt) => opt.trim());
            }
        };
    }

    // Mobile menu toggle
    const mobileMenuToggle = document.getElementById(
        "mobile-menu-toggle"
    );
    const mainNav = document.querySelector(".main-nav ul");

    if (mobileMenuToggle && mainNav) {
        mobileMenuToggle.addEventListener("click", () => {
            mainNav.classList.toggle("show");
        });
    }

    // Add active class to current page in navigation
    const currentPath = window.location.pathname;
    const navLinks =
        document.querySelectorAll(".main-nav a");

    navLinks.forEach((link) => {
        if (
            link.getAttribute("href") === currentPath ||
            (currentPath.includes(
                link.getAttribute("href")
            ) &&
                link.getAttribute("href") !== "/")
        ) {
            link.parentElement.classList.add("active");
        }
    });

    // Sidebar active links
    const sidebarLinks = document.querySelectorAll(
        ".docs-sidebar-link"
    );

    sidebarLinks.forEach((link) => {
        if (
            link.getAttribute("href") === currentPath ||
            currentPath === link.getAttribute("href")
        ) {
            link.classList.add("active");
        }
    });
});

/**
 * Initialize dropdown menus
 */
function initDropdowns() {
    const dropdowns =
        document.querySelectorAll(".dropdown");

    dropdowns.forEach((dropdown) => {
        // Add hover event listeners if needed
        dropdown.addEventListener(
            "mouseenter",
            function () {
                this.querySelector(
                    ".dropdown-menu"
                )?.classList.add("active");
            }
        );

        dropdown.addEventListener(
            "mouseleave",
            function () {
                this.querySelector(
                    ".dropdown-menu"
                )?.classList.remove("active");
            }
        );

        // For touch devices
        dropdown.addEventListener("click", function (e) {
            if (window.innerWidth <= 768) {
                const menu = this.querySelector(
                    ".dropdown-menu"
                );
                if (menu) {
                    e.preventDefault();
                    menu.classList.toggle("active");
                }
            }
        });
    });
}

/**
 * Utility function to toggle element visibility
 */
function toggleVisibility(element) {
    if (element) {
        element.classList.toggle("hidden");
    }
}

/**
 * Utility function for search in content
 */
function searchContent(
    searchText,
    items,
    getSearchableText,
    displayItem
) {
    searchText = searchText.toLowerCase();

    items.forEach((item) => {
        const text = getSearchableText(item).toLowerCase();
        const visible = text.includes(searchText);
        displayItem(item, visible);
    });
}

/**
 * Detect the user's preferred language and set up localization
 */
function setupLocalization() {
    const userLang =
        navigator.language || navigator.userLanguage;
    document.documentElement.setAttribute("lang", userLang);
}

// Export functions for use in other scripts
window.captoc = {
    toggleVisibility,
    searchContent,
    parseOptions: window.parseOptions,
};
