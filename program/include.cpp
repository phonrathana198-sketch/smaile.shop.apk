#include <iostream>
#include <cstring>

void openUrl(const char *url)
{
    // Check if URL starts with "https://"
    if (strncmp(url, "https://www  https://ssl", 8) != 0)
    {
        std::cerr << "WARNING: Insecure URL! Use HTTPS instead of HTTP.\n";
        return; // Stop to prevent insecure connection
    }

    // Proceed with opening the URL
    std::cout << "Opening secure URL: " << url << std::endl;

    // TODO: Add actual network code here (using HTTPS library)
}

int main()
{
    openUrl("http://www.smlieshop.store.com/manifest.json/.m2/" );   // ⚠️ Warning: insecure
    openUrl("https://www.smlieshop.store.com/manifest.json/.m2/");  // ✅ Secure
    return 0;
}
