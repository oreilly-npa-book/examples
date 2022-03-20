from lxml import etree
xslRoot = etree.fromstring(bytes(open("template.xslt").read(), encoding='utf8'))
transform = etree.XSLT(xslRoot)

xmlRoot = etree.fromstring(bytes(open("data.xml").read(), encoding='utf8'))
transRoot = transform(xmlRoot)

print(transRoot)