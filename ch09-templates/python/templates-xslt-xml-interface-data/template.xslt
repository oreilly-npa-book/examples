<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:template match="/">
        <xsl:for-each select="interfaces/interface">
        interface <xsl:value-of select="name"/>
            ip address <xsl:value-of select="ipv4addr"/>
        </xsl:for-each>
    </xsl:template>
</xsl:stylesheet>