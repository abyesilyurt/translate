## Translate CLI

This is a simple CLI tool to translate PDFs using the Google Translate API.
Especially, it is useful for translating PDFs that are not text-based, such as scanned documents.

## Usage

```bash
translate nl path/to/input.pdf
```

This command will translate the input PDF to Dutch and save the output as `path/to/input.pdf.translated.nl.pdf`.

You can omit the language code to translate to English by default.

```bash
translate path/to/input.pdf
```
