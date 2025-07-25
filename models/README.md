# MITIE Spanish Language Models

This directory should contain the MITIE Spanish language model files for Named Entity Recognition.

## Required Files

The application expects these files:
- `ner_model.dat` - Main NER model file (~450MB)
- `total_word_feature_extractor.dat` - Feature extractor

## Download Instructions

### Option 1: Use Makefile (Recommended)
```bash
make download-model
```

### Option 2: Manual Download
```bash
wget https://sourceforge.net/projects/mitie.mirror/files/v0.4/MITIE-models-v0.2-Spanish.zip/download -O models/spanish_model.zip
unzip models/spanish_model.zip -d models/
mv models/MITIE-models/spanish/* models/
rm -rf models/MITIE-models models/spanish_model.zip
```

### Option 3: Alternative Sources
You can also obtain Spanish MITIE models from:
- [MITIE official repository](https://github.com/mit-nlp/MITIE)
- Train your own using MITIE training tools

## Note

Model files are excluded from git (via .gitignore) due to their large size (~450MB).
Download them separately after cloning the repository.