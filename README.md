### **_Reume_** *(noun)*
**/ˈrɛu̯m(ə)/**

1. Obsolete form of realm
2. An anagram of resume (when pluralised)

[[src]](https://en.wiktionary.org/wiki/reume)

---

Reumes is a resume builder that can turn JSON Resume compliant data (in JSON or YAML format)
into anything that can be generated from code.

### Usage

Reumes essentially does three things, in this order:
1. Read in and validate resume data
2. Use that data to fill in a provided or custom [jinja](https://jinja.palletsprojects.com/en/stable/) template
3. Optionally postprocess the template using another program (such as [typst](https://typst.app/))

### Getting Started
- To learn how to make valid resume data, see the [JSON Resume Schema](https://jsonresume.org/schema)
- For example templates, see [the templates directory](./resume/templates).

---

This was made for my personal use, but I will do my best to support others' needs if there is demand.
