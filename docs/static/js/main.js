/**
 * CAP TOC - Main JavaScript
 * Modern UI interactions and utilities
 */

document.addEventListener("DOMContentLoaded", () => {
    // Initialize UI elements
    initThemeToggle();
    initMobileMenu();
    initSidebar();
    initDropdowns();
    setupActiveLinks();
    setupGenericContentInteractions();

    // Register utility functions globally
    window.captoc = {
        toggleVisibility,
        searchContent,
        parseOptions,
    };
});

/**
 * Initialize theme toggle (light/dark mode)
 */
function initThemeToggle() {
    const themeToggleBtn = document.getElementById(
        "theme-toggle-btn"
    );
    const prefersDarkScheme = window.matchMedia(
        "(prefers-color-scheme: dark)"
    );

    // Set initial theme based on user preference or saved setting
    const savedTheme = localStorage.getItem("theme");

    if (savedTheme) {
        document.documentElement.setAttribute(
            "data-theme",
            savedTheme
        );
    } else if (prefersDarkScheme.matches) {
        document.documentElement.setAttribute(
            "data-theme",
            "dark"
        );
    } else {
        document.documentElement.setAttribute(
            "data-theme",
            "light"
        );
    }

    // Handle theme toggle click
    if (themeToggleBtn) {
        themeToggleBtn.addEventListener("click", () => {
            const currentTheme =
                document.documentElement.getAttribute(
                    "data-theme"
                ) || "light";
            const newTheme =
                currentTheme === "light" ? "dark" : "light";

            document.documentElement.setAttribute(
                "data-theme",
                newTheme
            );
            localStorage.setItem("theme", newTheme);
        });
    }
}

/**
 * Initialize mobile menu toggle
 */
function initMobileMenu() {
    const mobileMenuToggle = document.getElementById(
        "mobile-menu-toggle"
    );
    const mainNav = document.querySelector(".main-nav");

    if (mobileMenuToggle && mainNav) {
        mobileMenuToggle.addEventListener("click", () => {
            mainNav.classList.toggle("show");

            // Toggle menu icon appearance
            const menuIcon =
                mobileMenuToggle.querySelector(
                    ".menu-icon"
                );
            if (menuIcon) {
                menuIcon.classList.toggle("active");
            }
        });

        // Close mobile menu when clicking outside
        document.addEventListener("click", (e) => {
            if (
                !e.target.closest(".main-nav") &&
                !e.target.closest("#mobile-menu-toggle") &&
                mainNav.classList.contains("show")
            ) {
                mainNav.classList.remove("show");
                const menuIcon =
                    mobileMenuToggle.querySelector(
                        ".menu-icon"
                    );
                if (menuIcon) {
                    menuIcon.classList.remove("active");
                }
            }
        });

        // Close mobile menu on ESC key
        document.addEventListener("keydown", (e) => {
            if (
                e.key === "Escape" &&
                mainNav.classList.contains("show")
            ) {
                mainNav.classList.remove("show");
                const menuIcon =
                    mobileMenuToggle.querySelector(
                        ".menu-icon"
                    );
                if (menuIcon) {
                    menuIcon.classList.remove("active");
                }
            }
        });
    }
}

/**
 * Initialize sidebar behavior
 */
function initSidebar() {
    const sidebar = document.querySelector(".sidebar");

    if (sidebar) {
        // Add automatic collapsing for mobile
        const toggleSidebarVisibility = () => {
            if (window.innerWidth <= 768) {
                sidebar.classList.add("collapsed");
            } else {
                sidebar.classList.remove("collapsed");
            }
        };

        // Run once on load
        toggleSidebarVisibility();

        // Run on resize
        window.addEventListener(
            "resize",
            toggleSidebarVisibility
        );
    }
}

/**
 * Initialize dropdown menus
 */
function initDropdowns() {
    const dropdowns =
        document.querySelectorAll(".dropdown");

    const isMobile = () => window.innerWidth <= 640;

    dropdowns.forEach((dropdown) => {
        const link = dropdown.querySelector(".nav-link");
        const menu = dropdown.querySelector(
            ".dropdown-menu"
        );

        if (link && menu) {
            link.addEventListener("click", (e) => {
                // On mobile, prevent default and toggle dropdown
                if (isMobile()) {
                    e.preventDefault();
                    menu.classList.toggle("active");

                    // Close other open dropdowns
                    dropdowns.forEach((otherDropdown) => {
                        if (otherDropdown !== dropdown) {
                            const otherMenu =
                                otherDropdown.querySelector(
                                    ".dropdown-menu"
                                );
                            if (
                                otherMenu &&
                                otherMenu.classList.contains(
                                    "active"
                                )
                            ) {
                                otherMenu.classList.remove(
                                    "active"
                                );
                            }
                        }
                    });
                }
            });
        }
    });

    // Close dropdowns when clicking outside
    document.addEventListener("click", (e) => {
        if (!e.target.closest(".dropdown")) {
            document
                .querySelectorAll(".dropdown-menu.active")
                .forEach((menu) => {
                    menu.classList.remove("active");
                });
        }
    });

    // Adjust dropdowns behavior when window resizes
    window.addEventListener("resize", () => {
        if (!isMobile()) {
            document
                .querySelectorAll(".dropdown-menu.active")
                .forEach((menu) => {
                    menu.classList.remove("active");
                });
        }
    });
}

/**
 * Set active links in navigation and sidebar
 */
function setupActiveLinks() {
    const currentPath = window.location.pathname;

    // Add active class to sidebar links
    document
        .querySelectorAll(".sidebar-link")
        .forEach((link) => {
            if (
                link.getAttribute("href") === currentPath ||
                currentPath.endsWith(
                    link.getAttribute("href")
                )
            ) {
                link.classList.add("active");

                // Ensure parent containers are visible
                const section = link.closest(
                    ".sidebar-section"
                );
                if (section) {
                    section.classList.add("active");
                }
            }
        });

    // Add active class to nav links
    document
        .querySelectorAll(".nav-link")
        .forEach((link) => {
            if (
                link.getAttribute("href") === currentPath ||
                (link.getAttribute("href") !== "/" &&
                    currentPath.includes(
                        link.getAttribute("href")
                    ))
            ) {
                link.classList.add("active");
                const navItem = link.closest(".nav-item");
                if (navItem) {
                    navItem.classList.add("active");
                }
            }
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
 * Parse options from string (for nguphap template)
 */
function parseOptions(optionsStr) {
    try {
        // Try to parse JSON string
        return JSON.parse(optionsStr.replace(/'/g, '"'));
    } catch (e) {
        // Fallback to simple comma separation if JSON parse fails
        return optionsStr
            .split(",")
            .map((opt) => opt.trim());
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
 * Detect user's preferred language
 */
function getPreferredLanguage() {
    return (
        navigator.language || navigator.userLanguage || "vi"
    );
}

/**
 * Set up interactions for generic content items
 */
function setupGenericContentInteractions() {
    const contentItems =
        document.querySelectorAll(".content-item");

    contentItems.forEach((item) => {
        // Add click event to expand/collapse content
        item.addEventListener("click", function (e) {
            // Don't toggle if clicking on a link or button
            if (
                e.target.tagName === "A" ||
                e.target.tagName === "BUTTON"
            ) {
                return;
            }

            this.classList.toggle("expanded");
        });
    });
}
