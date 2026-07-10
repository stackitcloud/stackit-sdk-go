#!/usr/bin/env bash

# Default behavior is to NOT fail on gorelease errors. Strict mode is intended for CI pipeline usage.
STRICT_MODE=0
REPORT_FILE="$(pwd)/gorelease_report.md"
BASE_BRANCH="origin/main"

# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --strict) STRICT_MODE=1 ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

# Initialize/clear the report file
echo -e "# gorelease report\n\ngenerated on $(date)" > "$REPORT_FILE"

echo "Checking modules under services/ for changes against $BASE_BRANCH..."

# Iterate over all go.mod files found exactly one level deep in services/
for mod_file in services/*/go.mod; do
    # Skip if no matches were found (handles the case where services/ is empty)
    [ -e "$mod_file" ] || continue
    
    mod_dir=$(dirname "$mod_file")
    
    # Check if there are any git changes in this specific directory compared to origin/main
    # --quiet implies --exit-code (returns 1 if there are differences)
    if ! git diff --quiet "$BASE_BRANCH" -- "$mod_dir"; then
        echo "Changes detected in $mod_dir. Running gorelease..."
        
        # Header for this module in the report
        echo -e "\n---\n\n### \`gorelease\` report for \`$mod_dir\`\n" >> "$REPORT_FILE"
        
        # Subshell to avoid directory path issues, though pushd/popd works too
        (
            cd "$mod_dir" || exit 1
            
            # Disable "exit on error" temporarily so we can capture the output
            # even if gorelease finds backwards incompatibilities (which causes exit code 1)
            set +e
            OUTPUT=$(gorelease 2>&1)
            EXIT_CODE=$?
            set -e
            
            # Append findings to the absolute path of the report file
            if [ -n "$OUTPUT" ]; then
                echo -e "\`\`\`\n$OUTPUT\n\`\`\`" >> "$REPORT_FILE"
            else
                echo "No API breaking changes or issues detected." >> "$REPORT_FILE"
            fi
            
            echo "" >> "$REPORT_FILE" # Blank line for spacing
            
            # Handle strict mode exit if gorelease failed (exit code != 0)
            if [ $EXIT_CODE -ne 0 ]; then
                echo "FAIL: gorelease found issues/failed for $mod_dir (Exit code: $EXIT_CODE)"
                
                if [ $STRICT_MODE -eq 1 ]; then
                    echo "Strict mode is enabled. Failing the script immediately."
                    exit $EXIT_CODE
                fi
            else
                echo "PASS: $mod_dir passed."
            fi
        )
        
        # If the subshell exited with an error (which happens in strict mode), 
        # propagate that exit code to the main script
        SUBSHELL_EXIT=$?
        if [ $STRICT_MODE -eq 1 ] && [ $SUBSHELL_EXIT -ne 0 ]; then
            exit $SUBSHELL_EXIT
        fi

    else
        echo "No changes in $mod_dir. Skipping."
    fi
done

echo ""
echo "Done. Aggregate report generated at: $REPORT_FILE"
