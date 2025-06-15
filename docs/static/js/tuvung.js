// Vocabulary (tuvung) page functionality
document.addEventListener("DOMContentLoaded", () => {
    // Pre-load all card backs to ensure proper rendering
    const cardBacks =
        document.querySelectorAll(".card-back");
    cardBacks.forEach((cardBack) => {
        // Force rendering to ensure content is visible when shown
        cardBack.style.visibility = "hidden";
        cardBack.classList.remove("hidden");
        setTimeout(() => {
            cardBack.classList.add("hidden");
            cardBack.style.visibility = "visible";
        }, 10);
    });

    // Vocabulary card functionality
    const cards = document.querySelectorAll(
        ".vocabulary-card"
    );
    const toggleAllBtn =
        document.getElementById("toggle-all");
    const hideAllBtn = document.getElementById("hide-all");
    const searchInput =
        document.getElementById("search-input");
    const searchButton =
        document.getElementById("search-button");

    // Set up each card
    cards.forEach((card) => {
        const toggleButton = card.querySelector(
            ".toggle-button"
        );
        const cardBack = card.querySelector(".card-back");

        // Toggle card functionality
        if (toggleButton) {
            toggleButton.addEventListener("click", () => {
                cardBack.classList.toggle("hidden");

                if (cardBack.classList.contains("hidden")) {
                    toggleButton.textContent = "Xem";
                } else {
                    toggleButton.textContent = "Ẩn";
                }
            });
        }
    });

    // Show all cards
    if (toggleAllBtn) {
        toggleAllBtn.addEventListener("click", () => {
            cards.forEach((card) => {
                const cardBack =
                    card.querySelector(".card-back");
                const toggleButton = card.querySelector(
                    ".toggle-button"
                );

                cardBack.classList.remove("hidden");
                toggleButton.textContent = "Ẩn";
            });
        });
    }

    // Hide all cards
    if (hideAllBtn) {
        hideAllBtn.addEventListener("click", () => {
            cards.forEach((card) => {
                const cardBack =
                    card.querySelector(".card-back");
                const toggleButton = card.querySelector(
                    ".toggle-button"
                );

                cardBack.classList.add("hidden");
                toggleButton.textContent = "Xem";
            });
        });
    }

    // Search functionality
    if (searchButton && searchInput) {
        searchButton.addEventListener("click", () => {
            performSearch();
        });

        searchInput.addEventListener("keypress", (e) => {
            if (e.key === "Enter") {
                performSearch();
            }
        });

        function performSearch() {
            const searchTerm = searchInput.value
                .toLowerCase()
                .trim();

            if (searchTerm === "") {
                // Show all cards if search is empty
                cards.forEach((card) => {
                    card.style.display = "block";
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
                } else {
                    card.style.display = "none";
                }
            });
        }
    }
});
