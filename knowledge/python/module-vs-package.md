---
id: python-module-vs-package-boundary
title_zh: module vs package
title_en: Python Module vs Package
type: true-false
body_format: bilingual-section
tags: [python, import, package]
difficulty: 2
question_zh: "在 Python 中，module 通常是一個檔案，而 package 通常是一組相關模組的目錄結構。"
question_en: "In Python, a module is usually a single file, while a package is usually a directory structure containing related modules."
answer: true
clickbait_zh: "你以為只是 import 不到，實際上你可能連邊界都分不清。"
clickbait_en: "If an import feels confusing, the real problem may be that the module-package boundary is still blurry."
review_hint_zh: "module 偏單檔，package 偏目錄與組織。"
review_hint_en: "A module is usually one file; a package is a directory-level grouping."
confusion_with: [python-import-module-loading]
metaphor_seed: [單檔, 目錄分組, 邊界]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

在 Python 裡，module 通常可以理解成一個 `.py` 檔案，裡面放函式、類別或變數。
package 則比較像一個用來組織多個 module 的資料夾結構，讓程式碼可以分層管理。
所以這兩者都和 import 有關，但層級並不一樣。

## en

In Python, a module is usually a single `.py` file containing functions, classes, or variables.
A package is more like a directory structure used to organize multiple related modules.
Both matter for imports, but they represent different levels of organization.
