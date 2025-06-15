// Grammar (nguphap) page functionality
document.addEventListener("DOMContentLoaded", () => {
    // Pre-load all answer results to ensure proper rendering
    const answerResults = document.querySelectorAll(
        ".answer-result"
    );
    answerResults.forEach((result) => {
        // Force rendering to ensure content is visible when shown
        result.style.visibility = "hidden";
        result.classList.remove("hidden");
        setTimeout(() => {
            result.classList.add("hidden");
            result.style.visibility = "visible";
        }, 10);
    });

    // Grammar question functionality
    const questions = document.querySelectorAll(
        ".grammar-question"
    );
    const showAllBtn = document.getElementById(
        "show-all-answers"
    );
    const hideAllBtn = document.getElementById(
        "hide-all-answers"
    );

    // Set up each question
    questions.forEach((question) => {
        const options = question.querySelectorAll(
            ".answer-option input"
        );
        const checkBtn = question.querySelector(
            ".check-answer-btn"
        );
        const resultDiv = question.querySelector(
            ".answer-result"
        );

        // Check answer button
        if (checkBtn) {
            checkBtn.addEventListener("click", () => {
                let selectedOption = null;

                options.forEach((option) => {
                    if (option.checked) {
                        selectedOption = option;
                    }
                });

                if (selectedOption) {
                    const parentOption =
                        selectedOption.closest(
                            ".answer-option"
                        );
                    const isCorrect =
                        parentOption.classList.contains(
                            "correct"
                        );

                    options.forEach((opt) => {
                        opt.disabled = true;
                    });

                    // Show all options correctly
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
                            } else if (
                                opt.querySelector("input")
                                    .checked
                            ) {
                                if (
                                    !opt.classList.contains(
                                        "correct"
                                    )
                                ) {
                                    opt.classList.add(
                                        "highlight-incorrect"
                                    );
                                }
                            }
                        });

                    // Show the result
                    resultDiv.classList.remove("hidden");
                    checkBtn.textContent = "Đã kiểm tra";
                    checkBtn.disabled = true;
                } else {
                    alert("Vui lòng chọn một đáp án");
                }
            });
        }
    });

    // Show all answers
    if (showAllBtn) {
        showAllBtn.addEventListener("click", () => {
            questions.forEach((question) => {
                const options = question.querySelectorAll(
                    ".answer-option input"
                );
                const resultDiv = question.querySelector(
                    ".answer-result"
                );
                const checkBtn = question.querySelector(
                    ".check-answer-btn"
                );

                // Disable all options
                options.forEach((opt) => {
                    opt.disabled = true;
                });

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
                        }
                    });

                // Show result
                resultDiv.classList.remove("hidden");
                checkBtn.textContent = "Đã kiểm tra";
                checkBtn.disabled = true;
            });
        });
    }

    // Hide all answers
    if (hideAllBtn) {
        hideAllBtn.addEventListener("click", () => {
            questions.forEach((question) => {
                const options = question.querySelectorAll(
                    ".answer-option input"
                );
                const resultDiv = question.querySelector(
                    ".answer-result"
                );
                const checkBtn = question.querySelector(
                    ".check-answer-btn"
                );

                // Enable all options and uncheck them
                options.forEach((opt) => {
                    opt.disabled = false;
                    opt.checked = false;
                });

                // Remove highlights
                question
                    .querySelectorAll(".answer-option")
                    .forEach((opt) => {
                        opt.classList.remove(
                            "highlight-correct",
                            "highlight-incorrect"
                        );
                    });

                // Hide result
                resultDiv.classList.add("hidden");
                checkBtn.textContent = "Kiểm tra đáp án";
                checkBtn.disabled = false;
            });
        });
    }
});
