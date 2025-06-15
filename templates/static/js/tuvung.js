/**
 * Vocabulary (tuvung) page functionality
 */
document.addEventListener("DOMContentLoaded", () => {
    initCards();
    initSearch();
    initCardControls();
});

/**
 * Initialize vocabulary cards
 */
function initCards() {
    const cards = document.querySelectorAll(
        ".vocabulary-card"
    );

    cards.forEach((card, index) => {
        const toggleButton = card.querySelector(
            ".toggle-button"
        );
        const closeButton =
            card.querySelector(".close-button");
        const cardBack = card.querySelector(".card-back");

        // Add animations with delay based on index for staggered effect
        card.style.animationDelay = `${index * 50}ms`;
        card.classList.add("animate-in");

        // Toggle card functionality
        if (toggleButton) {
            toggleButton.addEventListener("click", () => {
                toggleCard(card);
            });
        }

        // Close button functionality
        if (closeButton) {
            closeButton.addEventListener("click", () => {
                closeCard(card);
            });
        }
    });
}

/**
 * Initialize search functionality
 */
function initSearch() {
    const searchInput =
        document.getElementById("search-input");

    if (searchInput) {
        const cards = document.querySelectorAll(
            ".vocabulary-card"
        );

        // Search as you type
        searchInput.addEventListener(
            "input",
            debounce(() => {
                const searchTerm = searchInput.value
                    .toLowerCase()
                    .trim();

                if (searchTerm === "") {
                    cards.forEach((card) => {
                        card.style.display = "block";
                        card.classList.add("animate-in");
                    });
                    return;
                }

                // Filter cards
                cards.forEach((card) => {
                    const kanji = card
                        .querySelector(".kanji")
                        .textContent.toLowerCase();
                    const reading = card
                        .querySelector(".reading")
                        .textContent.toLowerCase();
                    const meaning = card
                        .querySelector(".meaning")
                        .textContent.toLowerCase();

                    if (
                        kanji.includes(searchTerm) ||
                        reading.includes(searchTerm) ||
                        meaning.includes(searchTerm)
                    ) {
                        card.style.display = "block";
                        card.classList.add("animate-in");
                    } else {
                        card.style.display = "none";
                        card.classList.remove("animate-in");
                    }
                });
            }, 300)
        );

        // Clear search on Escape key
        searchInput.addEventListener("keydown", (e) => {
            if (e.key === "Escape") {
                searchInput.value = "";
                cards.forEach((card) => {
                    card.style.display = "block";
                    card.classList.add("animate-in");
                });
                searchInput.blur();
            }
        });
    }
}

/**
 * Initialize card control buttons
 */
function initCardControls() {
    const toggleAllBtn =
        document.getElementById("toggle-all");
    const hideAllBtn = document.getElementById("hide-all");
    const cards = document.querySelectorAll(
        ".vocabulary-card"
    );

    // Show all cards
    if (toggleAllBtn) {
        toggleAllBtn.addEventListener("click", () => {
            cards.forEach((card) => {
                openCard(card);
            });
        });
    }

    // Hide all cards
    if (hideAllBtn) {
        hideAllBtn.addEventListener("click", () => {
            cards.forEach((card) => {
                closeCard(card);
            });
        });
    }
}

/**
 * Toggle card open/closed state
 * @param {HTMLElement} card - The vocabulary card element
 */
function toggleCard(card) {
    const cardBack = card.querySelector(".card-back");
    const toggleButton = card.querySelector(
        ".toggle-button"
    );
    const toggleText =
        toggleButton.querySelector(".toggle-text");

    if (cardBack.classList.contains("hidden")) {
        openCard(card);
    } else {
        closeCard(card);
    }
}

/**
 * Open a vocabulary card
 * @param {HTMLElement} card - The vocabulary card element
 */
function openCard(card) {
    const cardBack = card.querySelector(".card-back");
    const toggleButton = card.querySelector(
        ".toggle-button"
    );
    const toggleText =
        toggleButton.querySelector(".toggle-text");

    cardBack.classList.remove("hidden");
    if (toggleText) toggleText.textContent = "Ẩn";

    // Add focus effect
    card.classList.add("card-focus");
}

/**
 * Close a vocabulary card
 * @param {HTMLElement} card - The vocabulary card element
 */
function closeCard(card) {
    const cardBack = card.querySelector(".card-back");
    const toggleButton = card.querySelector(
        ".toggle-button"
    );
    const toggleText =
        toggleButton.querySelector(".toggle-text");

    cardBack.classList.add("hidden");
    if (toggleText) toggleText.textContent = "Xem";

    // Remove focus effect
    card.classList.remove("card-focus");
}

/**
 * Debounce function to limit how often a function is called
 * @param {Function} func - The function to debounce
 * @param {number} wait - Wait time in milliseconds
 * @returns {Function} - Debounced function
 */
function debounce(func, wait = 300) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}
