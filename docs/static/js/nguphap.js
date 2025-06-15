/**
 * Grammar (nguphap) page functionality
 */
document.addEventListener("DOMContentLoaded", () => {
    initQuestions();
    initControlButtons();
});

/**
 * Initialize grammar quiz questions
 */
function initQuestions() {
    const questions = document.querySelectorAll(
        ".grammar-question"
    );

    questions.forEach((question, index) => {
        // Add animation delay for staggered appearance
        question.style.animationDelay = `${index * 100}ms`;
        question.classList.add("animate-in");

        // Handle option selection by clicking the entire answer option div
        question
            .querySelectorAll(".answer-option")
            .forEach((option) => {
                option.addEventListener(
                    "click",
                    function () {
                        // Skip if already answered
                        if (
                            question.classList.contains(
                                "answered"
                            )
                        )
                            return;

                        // Select the option
                        const radio =
                            this.querySelector(
                                ".option-radio"
                            );
                        if (radio && !radio.disabled) {
                            // Uncheck all other options
                            question
                                .querySelectorAll(
                                    ".option-radio"
                                )
                                .forEach((r) => {
                                    r.checked = false;
                                });

                            // Check this one
                            radio.checked = true;

                            // Update visual selection
                            question
                                .querySelectorAll(
                                    ".answer-option"
                                )
                                .forEach((opt) => {
                                    opt.classList.remove(
                                        "selected"
                                    );
                                });

                            this.classList.add("selected");
                        }
                    }
                );
            });

        // Check answer button
        const checkBtn = question.querySelector(
            ".check-answer-btn"
        );
        if (checkBtn) {
            checkBtn.addEventListener("click", () => {
                checkAnswer(question);
            });
        }
    });
}

/**
 * Check answer for a question
 * @param {HTMLElement} question - The question container
 */
function checkAnswer(question) {
    const radio = question.querySelector(
        ".option-radio:checked"
    );

    if (radio) {
        // Mark question as answered
        question.classList.add("answered");

        // Get selected option element
        const selectedOption = radio.closest(
            ".answer-option"
        );
        const isCorrect =
            selectedOption.classList.contains("correct");

        // Disable all options
        question
            .querySelectorAll(".option-radio")
            .forEach((opt) => {
                opt.disabled = true;
            });

        // Show correct/incorrect highlights
        question
            .querySelectorAll(".answer-option")
            .forEach((opt) => {
                if (opt.classList.contains("correct")) {
                    opt.classList.add("highlight-correct");
                } else if (
                    opt.querySelector(".option-radio")
                        .checked
                ) {
                    if (
                        !opt.classList.contains("correct")
                    ) {
                        opt.classList.add(
                            "highlight-incorrect"
                        );
                    }
                }
            });

        // Show result
        const resultDiv = question.querySelector(
            ".answer-result"
        );
        resultDiv.classList.remove("hidden");
        resultDiv.classList.add("fade-in");

        // Update button
        const checkBtn = question.querySelector(
            ".check-answer-btn"
        );
        checkBtn.textContent = "Đã kiểm tra";
        checkBtn.disabled = true;

        // Add result class
        if (isCorrect) {
            question.classList.add("answered-correct");
        } else {
            question.classList.add("answered-incorrect");
        }
    } else {
        // Show alert if nothing selected
        const alert = document.createElement("div");
        alert.className = "alert alert-warning";
        alert.textContent = "Vui lòng chọn một đáp án";

        // Remove existing alerts
        const existingAlert =
            question.querySelector(".alert");
        if (existingAlert) {
            question.removeChild(existingAlert);
        }

        question.appendChild(alert);
        setTimeout(() => {
            alert.classList.add("show");
            setTimeout(() => {
                alert.classList.remove("show");
                setTimeout(() => {
                    if (question.contains(alert)) {
                        question.removeChild(alert);
                    }
                }, 300);
            }, 2000);
        }, 10);
    }
}

/**
 * Initialize control buttons for showing/hiding answers
 */
function initControlButtons() {
    const showAllBtn = document.getElementById(
        "show-all-answers"
    );
    const hideAllBtn = document.getElementById(
        "hide-all-answers"
    );
    const questions = document.querySelectorAll(
        ".grammar-question"
    );

    // Show all answers
    if (showAllBtn) {
        showAllBtn.addEventListener("click", () => {
            questions.forEach((question) => {
                // Skip already answered questions
                if (question.classList.contains("answered"))
                    return;

                // Highlight correct answers
                question
                    .querySelectorAll(".answer-option")
                    .forEach((opt) => {
                        if (
                            opt.classList.contains(
                                "correct"
                            )
                        ) {
                            opt.classList.add(
                                "highlight-correct"
                            );

                            // Check the correct radio button
                            const radio =
                                opt.querySelector(
                                    ".option-radio"
                                );
                            if (radio) radio.checked = true;
                        }
                    });

                // Show result
                const resultDiv = question.querySelector(
                    ".answer-result"
                );
                resultDiv.classList.remove("hidden");

                // Disable options
                question
                    .querySelectorAll(".option-radio")
                    .forEach((opt) => {
                        opt.disabled = true;
                    });

                // Update button
                const checkBtn = question.querySelector(
                    ".check-answer-btn"
                );
                checkBtn.textContent = "Đã kiểm tra";
                checkBtn.disabled = true;

                // Mark as answered
                question.classList.add("answered");
            });
        });
    }

    // Hide all answers
    if (hideAllBtn) {
        hideAllBtn.addEventListener("click", () => {
            questions.forEach((question) => {
                // Reset options
                question
                    .querySelectorAll(".option-radio")
                    .forEach((opt) => {
                        opt.disabled = false;
                        opt.checked = false;
                    });

                // Clear highlights
                question
                    .querySelectorAll(".answer-option")
                    .forEach((opt) => {
                        opt.classList.remove(
                            "highlight-correct",
                            "highlight-incorrect",
                            "selected"
                        );
                    });

                // Hide result
                const resultDiv = question.querySelector(
                    ".answer-result"
                );
                resultDiv.classList.add("hidden");

                // Reset button
                const checkBtn = question.querySelector(
                    ".check-answer-btn"
                );
                checkBtn.textContent = "Kiểm tra đáp án";
                checkBtn.disabled = false;

                // Remove answered state
                question.classList.remove(
                    "answered",
                    "answered-correct",
                    "answered-incorrect"
                );
            });
        });
    }
}

/**
 * Check if an element is currently visible in the viewport
 * @param {HTMLElement} el - The element to check
 * @returns {boolean} - True if element is in viewport
 */
function isElementInViewport(el) {
    const rect = el.getBoundingClientRect();
    return (
        rect.top >= 0 &&
        rect.left >= 0 &&
        rect.bottom <=
            (window.innerHeight ||
                document.documentElement.clientHeight) &&
        rect.right <=
            (window.innerWidth ||
                document.documentElement.clientWidth)
    );
}
