import { Injectable } from '@angular/core';
import { jsPDF } from 'jspdf';
import { formatDate } from '@angular/common';
import { RelatorioView } from './relatorio-view.inteface';

@Injectable()
export class RelatorioViewComponent {

  public row = 0;
  private doc: jsPDF;

  constructor() {
  }

  private header(doc, relatorio: RelatorioView) {
    //doc.height(relatorio.title.pageHeight);
    doc.setFont(relatorio.title.font);
    doc.setFontStyle(relatorio.title.style);
    doc.setFontSize(relatorio.title.size);
    doc.setTextColor(relatorio.title.color[0], relatorio.title.color[1], relatorio.title.color[2]);
    doc.text(relatorio.title.name, relatorio.title.position, relatorio.title.row);
    // ResetaDados
    doc.setFont(relatorio.title.font);
    doc.setFontStyle('normal');
    doc.setFontSize(10);
    this.row = relatorio.title.row + (relatorio.title.size / 8);

    // TituloBody
    relatorio.fields.forEach(field => {
      doc.setFont(field.font);
      doc.setFontStyle(field.styleLabel);
      doc.setFontSize(field.size);
      doc.setTextColor(field.color[0], field.color[1], field.color[2]);
      doc.text(field.label, field.position, field.row);
      this.row = field.row + (field.size / 5);
    });
  }

  // Body
  private body(doc, relatorio, row) {
    row += 2;
    const rowInicial = row;
    let maxSize = 0;
    let rows = 1;
    let total = [];

    relatorio.dataSet.forEach((data, item) => {
      relatorio.fields.forEach(field => {
        doc.setFont(field.font);
        doc.setFontStyle(field.style);
        doc.setFontSize(field.size);
        doc.setTextColor(field.color[0], field.color[1], field.color[2]);

        let valueform = String(data[field.name] ? data[field.name] : '');
        if (field.format) {
          if (String(field.format).indexOf('0.') >= 0) {
            const digits = String(field.format).substring(String(field.format).indexOf('0.') + 2).length;
            valueform = Number(data[field.name] ? data[field.name] : '0').toLocaleString('pt-BR', { minimumFractionDigits: digits });
          } else if (data[field.name]) {
            valueform = formatDate(data[field.name], field.format, 'en');
          }
        }
        doc.text(valueform, field.position, row);
        maxSize = maxSize > field.size ? maxSize : field.size;
        if (field.totalize === true) {
          let value = total.filter(a => a.field === field.name)[0];
          if (value) {
            value['value'] += Number.parseFloat(valueform);
          } else {
            total.push({
              field: field.name,
              value: data[field.name],
              font: field.font,
              style: field.style,
              styleLabel: field.styleLabel,
              size: field.size,
              color: field.color,
              format: field.format,
              position: field.position
            });
          }

        }
      });
      row += 2 + (maxSize / 5);
      rows += 1;
      if ((rows - 1) === relatorio.title.rowPage) {
        if (row <= 292) {
          this.footer(this.doc, relatorio);
        }
        rows = 1;
        if ((relatorio.dataSet.length - 1) > item) {
          doc.addPage();
          this.header(doc, relatorio);
        }
        row = rowInicial;
      }
    });

    if (total.length > 0) {
      total.forEach(field => {
        doc.setFont(field.font);
        doc.setFontStyle(field.styleLabel);
        doc.setFontSize(field.size);
        doc.setTextColor(field.color[0], field.color[1], field.color[2]);
        let valueform = field.value.toString();
        if (field.format) {
          valueform = field.value.toLocaleString('pt-BR', { minimumFractionDigits: 2 });
        }

        doc.text(valueform, field.position, row);

        maxSize = maxSize > field.size ? maxSize : field.size;
      });

      row += 2 + (maxSize / 5);
      if (row > 292) {
        rows = 1;
        doc.addPage();
        this.header(doc, relatorio);
        row = rowInicial;
      }
      this.footer(this.doc, relatorio);
    }

  }

  private footer(doc, relatorio) {
    doc.setFont('Courier');
    doc.setFontStyle('Bold');
    doc.setFontSize(9);
    doc.text('___________________________________________________________________________________________________________________ Coneplus ®', relatorio.title.margin, relatorio.title.pageHeight - relatorio.title.margin);
  }

  public imprimir(relatorio, pdf) {
    this.doc = new jsPDF({
      orientation: relatorio.orientation ? relatorio.orientation : "portrait",
      unit: relatorio.unit ? relatorio.unit : "cm",
      format: relatorio.format ? relatorio.format : [4, 2],
    });

    // Optional - set properties on the document
    this.doc.setProperties({
      title: relatorio.modelo,
      subject: '',
      author: 'DebtConvert Sistemas',
      keywords: '',
      creator: 'Coneplus'
    });

    this.header(this.doc, relatorio);
    this.body(this.doc, relatorio, this.row);

    if (pdf) {
      this.doc.save(relatorio.modelo.toString() + formatDate(new Date(), 'yyyyMMdd', 'en') + '.pdf');
    } else {
      this.doc.output('dataurlnewwindow');
    }

  }

}
